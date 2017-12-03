package main

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/testdata/59-rename-method/dep"
)

func main() {
	println(macro.Raw("typeof dep.test", dep.Test))
	println(dep.Test())
}
