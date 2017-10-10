package main

import "errors"

func main() {
	err := errors.New("new error")
	panic(err)
}
