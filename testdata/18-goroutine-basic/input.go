package main

func main() {
	ch := make(chan string)
	go func() {
		ch <- "hi"
	}()

	println(<-ch)
}
