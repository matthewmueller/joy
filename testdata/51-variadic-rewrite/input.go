package main

import "github.com/matthewmueller/golly/js"

var one = "1"

func main() {
	println(1, test("a", "b", "c"))
	println(2, test("a"))
}

func test(two string, rest ...string) string {
	js.Rewrite("$3.map(function(a) { return $1 + $2 + a }).join(' ')", one, two, rest)
	return ""
}
