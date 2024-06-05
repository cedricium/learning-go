package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var input = `
{
	"created_at": "Mon Jun 05 00:00:01 -0700 2023"
}
`

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
	if err != nil {
		return err
	}
	*t = Timestamp(v)
	return nil
}

func main() {
	// var val map[string]interface{}
	// var val map[string]any       // 1.
	// var val map[string]time.Time // 2.
	var val map[string]Timestamp // 3.

	if err := json.Unmarshal([]byte(input), &val); err != nil {
		panic(err)
	}

	fmt.Println(val)
	for k, v := range val {
		fmt.Println(k, reflect.TypeOf(v))
	}
	fmt.Println(time.Time(val["created_at"]))
}

/*
1. Using `map[string]any` type, though difficult to compare times using strings
*/
// map[created_at:Mon Jun 5 00:00:01 +0000 2023]
// created_at string

/*
2. Using `map[string]time.Time` type results in the following error:
*/
// panic: parsing time "Mon Jun 5 00:00:01 +0000 2023" as "2006-01-02T15:04:05Z07:00": cannot parse "Mon Jun 5 00:00:01 +0000 2023" as "2006"

/*
3. Using `map[string]Timestamp` type, where `Timestamp` satisfies `json.Unmarshaler` interface
*/
// map[created_at:{0 63821545201 0x101059920}]
// created_at main.Timestamp
// 2023-06-05 00:00:01 -0700 PDT
