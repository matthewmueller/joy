package main

import (
	"github.com/matthewmueller/joy/js"
	"github.com/matthewmueller/joy/testdata/59-rename-method/dep"
)

func main() {
	println(js.Raw("typeof dep.test", dep.Test))
	println(dep.Test())
}
