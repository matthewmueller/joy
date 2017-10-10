package main

import "errors"

func main() {
	err := errors.New("new error")
	println(err.Error())
	panic(err)
}
