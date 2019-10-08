package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

// Decode is the example of using json.Decoder to decode JSON stream.
// To multiple JSON values:
// 1. For JSON numbers, append "\n" to the end of JSON value as dilimiter, it can be recognized by decoder.
// 2. For JSON strings / objects / arrays, "",{},[] works well, no need to add "\n".
func Decode() {
	var err error
	var n int
	var buf bytes.Buffer
	var str = `
[
  {"brand":"Intel","model":"i7-9700K"},
  {"brand":"AMD","model":"3900X"}
]
`

	// datas contains multiple JSON values.
	datas := [][]byte{
		// JSON number, use `\n` to split other number value.
		[]byte("123\n"),
		// another JSON number
		[]byte("3.1415926"),
		// JSON string
		[]byte(`"Hello, World!"`),
		// another JSON string
		[]byte(`"JSON example"`),
		// JSON object
		[]byte(`{"name":"Frank","skills":["go","c"]}`),
		// JSON array
		[]byte(str),
	}

	// Write bytes to bytes.Buffer(io.Writer).
	for _, data := range datas {
		if n, err = buf.Write(data); err != nil {
			fmt.Printf("Write() error: %v\n", err)
			return
		}
		fmt.Printf("Write() OK. n: %v, appended data: %s\n", n, string(data))
	}

	// Create a JSON decoder using bytes.Buffer(io.Reader)
	dec := json.NewDecoder(&buf)

	for {
		// Use empty interface to store decoded value.
		var v interface{}

		// Read from bytes.Buffer and decode JSON value one by one.
		if err = dec.Decode(&v); err != nil && err != io.EOF {
			fmt.Printf("Decode() error: %v\n", err)
			return
		} else if err == io.EOF {
			break
		}

		fmt.Printf("Decode() OK. type: %T, v: %v\n", v, v)
	}

	// Output:
	//Write() OK. n: 4, appended data: 123

	//Write() OK. n: 9, appended data: 3.1415926
	//Write() OK. n: 15, appended data: "Hello, World!"
	//Write() OK. n: 14, appended data: "JSON example"
	//Write() OK. n: 36, appended data: {"name":"Frank","skills":["go","c"]}
	//Write() OK. n: 79, appended data:
	//[
	//  {"brand":"Intel","model":"i7-9700K"},
	//  {"brand":"AMD","model":"3900X"}
	//]

	//Decode() OK. type: float64, v: 123
	//Decode() OK. type: float64, v: 3.1415926
	//Decode() OK. type: string, v: Hello, World!
	//Decode() OK. type: string, v: JSON example
	//Decode() OK. type: map[string]interface {}, v: map[name:Frank skills:[go c]]
	//Decode() OK. type: []interface {}, v: [map[brand:Intel model:i7-9700K] map[brand:AMD model:3900X]]
}

func main() {
	Decode()
}
