package main

import (
	"github.com/matthewmueller/golly/js"
)

// FetchOptions struct
type FetchOptions struct {
	Method string `js:"method"`
}

// Fetch fn
func Fetch(url string, options FetchOptions) (*Response, error) {
	fetch := js.RawFile("./unfetch.js")
	// res := make(chan Response, 1)
	println(fetch)
	// println(res)

	// js.Promise("window.fetch(url, options)").Then(res).Catch(err)
	// return res.Ok, err
	return nil, nil
}

// Response struct
type Response struct {
	responseText string
	Ok           bool
	Status       int
	StatusText   string
	URL          string
}

// JSON fn
// func (r *Response) JSON(v interface{}) (err error) {
// 	js.Promise(".json()").Then(v).Catch(err)
// 	return nil
// }

func main() {
	response, err := Fetch("http://google.com", FetchOptions{
		Method: "POST",
	})
	if err != nil {
		println(err)
		return
	}

	println(response)

	// var raw []byte
	// err = response.JSON(raw)
	// if err != nil {
	// 	println(err)
	// 	return
	// }

	// println(raw)
}
