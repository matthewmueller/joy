package main

func main() {
	test("a", "b", "c")
}

func test(results ...string) {
	println(results)
	test2("hi", results...)

	func(prefix string, results ...string) {
		// println(prefix, results)
	}("hello", results...)
}

func test2(prefix string, results ...string) {
	println(prefix, results)
}
