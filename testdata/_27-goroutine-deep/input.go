package main

func main() {
	ch := make(chan string)
	go func() {
		func() {
			go test(ch)
		}()
	}()

	println(<-ch)
	println(another())
}

func test(ch chan string) {
	test2(ch)
}

func test2(ch chan string) {
	ch <- "hi"
}

func another() string {
	return "another"
}
