/*
FetchAll fetches URLs in parallel and reports their times and sizes.
It can either accept a list of URLs as args or a file containing URLs.

Usage:

	fetchAll [URL…] [options]
	fetchAll --file FILE [options]

Options:

	URL…
		List of URLs to process, separated by a space.

	-f, --file FILE
		Read URLs from FILE, one per line.

	-o, --output FILE
		Output file. Defaults to stdout if not provided.

Examples:

	$ ./fetchAll https://google.com example.com gopl.io
	$ ./fetchAll -f "alexa-top-1000.txt" -o results.txt
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const KiB = 1024.0

var flagInputFile string
var flagOutputFile string

func init() {
	const (
		defaultFile = ""
		usageFile   = "the file containing URLs (one per line) to process"
		usageOutput = "the file to save output to, default stdout"
	)
	flag.StringVar(&flagInputFile, "file", defaultFile, usageFile)
	flag.StringVar(&flagInputFile, "f", defaultFile, usageFile+" (shorthand)")
	flag.StringVar(&flagOutputFile, "output", defaultFile, usageOutput)
	flag.StringVar(&flagOutputFile, "o", defaultFile, usageOutput+" (shorthand)")
}

func main() {
	flag.Parse()
	start := time.Now()
	ch := make(chan string)
	urls := make([]string, 0)

	var output *os.File
	if flagOutputFile == "" {
		output = os.Stdout
	} else {
		f, err := os.Create(flagOutputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetchAll: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		output = f
	}

	if flagInputFile == "" {
		urls = append(urls, os.Args[1:]...)
	} else {
		f, err := os.Open(flagInputFile)
		if err != nil {
			fmt.Fprintf(output, "fetchAll: %v\n", err)
			os.Exit(1)
		}
		defer f.Close()
		input := bufio.NewScanner(f)
		for input.Scan() {
			url := input.Text()
			urls = append(urls, url)
		}
	}

	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
		fmt.Fprintln(output, <-ch)
	}

	fmt.Fprintf(output, "\n%5.2fs elapsed in total\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()
	written, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%5.2fs\t%8.2FKiB\t%s", secs, float64(written)/KiB, url)
}
