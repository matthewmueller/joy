//go:generate go run main.go

package main

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/dom/curl"
	"github.com/matthewmueller/golly/internal/dom/parser"
)

var browserAPIURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/browser.webidl.xml"
var browserDocsURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/comments.json"

func generate() error {
	xml, err := curl.XML(browserAPIURL)
	if err != nil {
		return err
	}

	comments, err := curl.JSON(browserDocsURL)
	if err != nil {
		return err
	}

	defs, err := parser.Parse(xml, comments)
	if err != nil {
		return err
	}

	for _, d := range defs {
		if d.Kind() != "INTERFACE" {
			continue
		}

		children, err := d.Children()
		if err != nil {
			return err
		}
		_ = children
		// log.Infof("got children %d", )
	}

	// attr := idx["Attr"]
	// log.Infof("attr %+v", attr.Children())

	return nil
}

func main() {
	if e := generate(); e != nil {
		log.WithError(e).Fatalf("error generating")
	}
}
