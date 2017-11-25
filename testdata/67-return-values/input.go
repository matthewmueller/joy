package main

type Test interface{}

func test() (a string, t Test) {
	return a, t
}

func main() {
	a, t := test()
	println(a, t)
}
