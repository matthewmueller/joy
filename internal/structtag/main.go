package main

import (
	"log"

	"github.com/matthewmueller/joy/internal/structtag/calc"
)

func main() {

	c := &calc.Calculator{Buffer: string("2+5")}
	c.Init()

	if err := c.Parse(); err != nil {
		log.Fatal(err)
	}

	c.PrintSyntaxTree()
}
