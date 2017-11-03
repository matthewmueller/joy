package main

import "github.com/matthewmueller/golly/js"

func main() {
	println(1, test("a", "b", "c"))
	println(2, test("a"))
}

func test(one string, rest ...string) string {
	js.Rewrite("$2.map(function(a) { return $1 + a }).join(' ')", one, rest)
	return ""
}
