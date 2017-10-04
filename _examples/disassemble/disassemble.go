package main

import (
	"io/ioutil"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/jsast/disassembler"
)

func main() {
	ch, e := ioutil.ReadFile("./channel.js")
	if e != nil {
		log.WithError(e).Fatalf("couldn't read file")
	}

	res, e := disassembler.Disassemble(string(ch))
	if e != nil {
		log.WithError(e).Fatalf("couldn't disassemble file")
	}

	log.Infof("got res %+v", res)

}
