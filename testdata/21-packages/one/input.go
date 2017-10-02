package main

import (
	"github.com/matthewmueller/golly/testdata/21-packages/dep"
)

type test struct {
	String string
}

func main() {
	t := test{String: "hi"}
	println(dep.Dep("one") + t.String)
}
