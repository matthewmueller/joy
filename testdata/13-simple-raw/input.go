package main

import "github.com/matthewmueller/golly/js"

func main() {
	x := 15
	js.Raw("x += 5")
	println(x)
}
