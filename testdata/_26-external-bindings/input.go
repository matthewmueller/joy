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
	_ = fetch
	ch := make(chan Response, 1)

	js.Raw(`
		fetch(url, options)
			.then(function (res) {
				return ch.Send(res)
			})
			.catch(function (err) {
				console.log(err)
			})
	`)

	res := <-ch
	return &res, nil
}

// Response interface
// js:"response,global"
type Response struct{}

// // Response struct
// type Response struct {
// 	responseText string
// 	Ok           bool
// 	Status       int
// 	StatusText   string
// 	URL          string
// }

// JSON fn
func (res *Response) JSON(v interface{}) (string, error) {
	ch := make(chan string, 1)
	js.Raw(`
		res.json()
			.then(function (str) {
				return ch.Send(str)
			})
			.catch(function (err) {
				console.log(err)
			})
	`)

	return <-ch, nil
}

func main() {
	response, err := Fetch("https://api.github.com/users/matthewmueller/repos", FetchOptions{
		Method: "GET",
	})
	if err != nil {
		println(err)
		return
	}

	str, err := response.JSON(nil)
	if err != nil {
		println(err)
		return
	}
	println(str)
	// println(response.responseText)

	// var raw []byte
	// err = response.JSON(raw)
	// if err != nil {
	// 	println(err)
	// 	return
	// }

	// println(raw)
}
