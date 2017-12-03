package main

import (
	"github.com/matthewmueller/joy/macro"
)

var fn = macro.RawFile("./fn.js")

func main() {
	println(test("hello world"))
}

func test(msg string) string {
	macro.Rewrite("$1($2)", fn, msg)
	return ""
}
