package main

import (
	dep2 "github.com/matthewmueller/golly/_examples/packaging/another/dep"
	"github.com/matthewmueller/golly/_examples/packaging/dep"
)

func main() {
	println("hi " + dep.Get() + " " + Side() + " " + dep2.AnotherGet())
}

func Test() string {
	return "another test"
}

func hi() string {
	return "hi"
}
