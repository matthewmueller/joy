package main

import "github.com/matthewmueller/golly/js"

func call() interface{} {
	fetch := js.RawFile(`./fetch.js`)
	return fetch
}

func main() {
	fetch := call()
	_ = fetch
	res := js.Raw(`fetch("http://google.com")`)
	println(res)
}
