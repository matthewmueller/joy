package main

import (
	"github.com/matthewmueller/golly/js"
)

var fn = js.RawFile("./fn.js")

func main() {
	println(test("hello world"))
}

func test(msg string) string {
	js.Rewrite("$1($2)", fn, msg)
	return ""
}
