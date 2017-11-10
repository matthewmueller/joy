package main

import (
	"github.com/matthewmueller/golly/js"
	"github.com/matthewmueller/golly/testdata/59-rename-method/dep"
)

func main() {
	println(js.Raw("typeof dep.test", dep.Test))
	println(dep.Test())
}
