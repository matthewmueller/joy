package main

func main() {
	test()
	t := tester{}
	t.test()
}

func test() {
	println("func")
}

type tester struct{}

func (t *tester) test() {
	println("method")
}
