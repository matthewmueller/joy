package main

import "github.com/matthewmueller/joy/macro"

func call() interface{} {
	fetch := macro.RawFile(`./fetch.js`)
	return fetch
}

func main() {
	fetch := call()
	_ = fetch
	res := macro.Raw(`fetch("http://google.com")`)
	println(res)
}
