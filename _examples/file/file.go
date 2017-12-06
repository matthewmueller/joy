package main

import (
	"github.com/matthewmueller/joy/macro"
)

func main() {
	test()
}

func test() {
	macro.Rewrite("$1()", macro.File("./test.js"))
}
