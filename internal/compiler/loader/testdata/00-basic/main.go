package main

import (
	"fmt"

	"github.com/matthewmueller/joy/macro"
)

func main() {
	macro.Raw("console.log")
	fmt.Println("hi world")
}
