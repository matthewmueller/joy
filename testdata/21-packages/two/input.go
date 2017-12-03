package main

import (
	"github.com/matthewmueller/joy/testdata/21-packages/dep"
)

func main() {
	println(dep.Dep("two") + dep.A + dep.Another())
}
