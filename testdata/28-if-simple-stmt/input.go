package main

import "errors"

func main() {
	if x, err := test(); err != nil {
		println(err.Error())
	} else {
		println(x)
	}
}

func test() (string, error) {
	return "", errors.New("oh noz")
}
