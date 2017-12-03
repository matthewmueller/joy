package main

import "github.com/matthewmueller/joy/macro"

func main() {
	x := 15
	macro.Raw("x += 5")
	println(x)
}
