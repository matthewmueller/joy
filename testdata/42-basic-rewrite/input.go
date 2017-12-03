package main

import "github.com/matthewmueller/joy/macro"

func main() {
	test("hi")
}

func test(msg string) {
	macro.Rewrite("console.log($1)", msg)
}
