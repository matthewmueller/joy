package main

import (
	"fmt"

	"github.com/matthewmueller/joy/api/run/testdata/00-packages/dep"
)

func main() {
	fmt.Println(dep.Dep("a"))
}
