package main

import (
	zarg "github.com/matthewmueller/joy/testdata/21-packages/dep"
)

type test struct {
	String string
}

func main() {
	t := test{String: "hi"}
	println(zarg.Dep("one") + t.String)
}
