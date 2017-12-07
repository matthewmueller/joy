package main

import (
	"github.com/apex/log"
	"github.com/matthewmueller/joy/internal/env"
)

func main() {
	log.Infof("env %#v", env.Get())
}
