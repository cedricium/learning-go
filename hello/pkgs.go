package main

import (
	"fmt"

	/* running `$ go mod tidy` can help add missing modules */
	/* running `$ go get <external link to package>` (e.g. `$ go get rsc.io/quote`) */
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Opt())
}
