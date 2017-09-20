package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

// Test struct
type Test struct {
	Name      string `json:"name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	gosrc, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.WithError(err).Fatalf("couldn't read from stdin")
	}

	fmt.Println(string(gosrc))
}
