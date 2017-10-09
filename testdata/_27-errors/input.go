package main

import "errors"

func main() {
	err := errors.New("test")
	println(err.Error())
}
