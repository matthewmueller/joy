package main

import (
	"fmt"
	"os"

	"github.com/apex/log"
	"github.com/matthewmueller/golly"
)

func main() {
	file := os.Args[1]

	js, err := golly.CompileFile(file)
	if err != nil {
		log.WithError(err).Fatal("unable to compile to js")
	}

	fmt.Println(js)
	// log.Infof("file %s", file)
	// // defer os.Stdin.Close()
	// reader := bufio.NewReader(os.Stdin)
	// text, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.WithError(err).Fatal("unable to read from stdin")
	// }

	// js, err := golly.Compile(text)
	// if err != nil {
	// 	log.WithError(err).Fatal("unable to compile to js")
	// }

	// fmt.Println(js)
}
