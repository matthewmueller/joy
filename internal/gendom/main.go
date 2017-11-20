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

	dom := path.Join("..", "dom", "browser.webidl.xml")
	src, err := ioutil.ReadFile(dom)
	if err != nil {
		log.WithError(err).Fatalf("error reading file")
	}

	files, err := generator.Generate(string(src))
	if err != nil {
		log.WithError(err).Fatalf("error generating dom")
	}

	for _, file := range files {
		outpath := path.Join("dom", file.Name)
		if e := os.MkdirAll(path.Dir(outpath), os.ModePerm); e != nil {
			log.WithError(e).Fatalf("error making directory")
		}
		if e := ioutil.WriteFile(outpath, []byte(file.Source), os.ModePerm); e != nil {
			log.WithError(e).Fatalf("error writing file")
		}
	}
}
