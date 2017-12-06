package main

func main() {
	ch := make(chan string, 1)
	select {
	case str := <-ch:
		println(str)
	default:
		println("default")
	}
}
