package main

import "time"

func main() {
	test()
}

func test() {
	defer func() {
		println("deferred")
	}()
	time.Sleep(100 * time.Millisecond)
	println("slept")
}
