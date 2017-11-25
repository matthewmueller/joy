package main

type Test interface{}

func test() (a string, b, c int, t Test) {
	return a, b, c, t
}

func main() {
	a, b, c, t := test()
	println(len(a), b, c, t)
}
