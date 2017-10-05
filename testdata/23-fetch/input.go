package main

import "github.com/matthewmueller/golly/window"

func main() {
	response, err := window.Fetch("http://google.com", window.FetchOptions{
		Method: "POST",
	})
	if err != nil {
		println(err)
		return
	}

	raw, err := response.JSON()
	if err != nil {
		println(err)
		return
	}

	println(raw)
}
