package main

import (
	"github.com/apex/log"
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
				ch.Send(new Response({
					responseText: res.responseText,
					statusText: res.statusText,
					status: res.status,
					ok: res.ok,
					url: url
				}))
			})
			.catch(function (err) {
				console.log(err)
			})
	`)

	res := <-ch
	log.Infof("here?")
	return &res, nil
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
func (r *Response) JSON(v interface{}) (json string, err error) {
	println(r.responseText)
	return r.responseText, nil
	// return json.Unmarshal([]byte(r.responseText), v)
}

func main() {
	response, err := Fetch("https://api.github.com/users/matthewmueller/repos", FetchOptions{
		Method: "POST",
	})
	if err != nil {
		println(err)
		return
	}
	log.Infof("here...")
	str, err := response.JSON(nil)
	if err != nil {
		println(err)
		return
	}

	log.Infof("result %s", str)
	// println(response.responseText)

	// var raw []byte
	// err = response.JSON(raw)
	// if err != nil {
	// 	println(err)
	// 	return
	// }

	// println(raw)
}
