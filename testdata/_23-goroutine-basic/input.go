package main

func main() {
	ch := make(chan string)
	go func(msg string) {
		ch <- msg
	}("hi")

	println(<-ch)
}
