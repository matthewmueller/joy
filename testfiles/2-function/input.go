package main

func main() {
	greet := hello("anki")
	println(greet)
}

func hello(name string) (out string) {
	return "hello " + name
}
