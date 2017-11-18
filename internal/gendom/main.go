//go:generate go run main.go

package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/internal/gendom/generator"
)

func main() {
	log.SetHandler(text.New(os.Stderr))

	dom := path.Join("data", "browser.webidl.xml")
	src, err := ioutil.ReadFile(dom)
	if err != nil {
		log.WithError(err).Fatalf("error reading file")
	}

	files, err := generator.Generate(string(src))
	if err != nil {
		log.WithError(err).Fatalf("error generating dom")
	}

	log.Infof("got files %+v", files)
	// log.Infof("data path %s", data)
}
