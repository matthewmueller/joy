package main

import (
	"github.com/matthewmueller/golly/js"
)

// FetchOptions struct
type FetchOptions struct {
	Method string `js:"method"`
}

// Fetch fn
func Fetch(url string, options FetchOptions) (res Response, err error) {
	js.Promise("window.fetch(url, options)").Then(res).Catch(err)
	return res.Ok, err
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
func (r *Response) JSON(v interface{}) (err error) {
	js.Promise(".json()").Then(v).Catch(err)
	return nil
}

func main() {
	response, err := Fetch("http://google.com", FetchOptions{
		Method: "POST",
	})
	if err != nil {
		println(err)
		return
	}

	var raw []byte
	err = response.JSON(raw)
	if err != nil {
		println(err)
		return
	}

	println(raw)
}
