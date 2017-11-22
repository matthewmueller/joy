//go:generate go run main.go

package main

import (
	"fmt"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/dom/curl"
	"github.com/matthewmueller/golly/internal/dom/defs"
	"github.com/matthewmueller/golly/internal/dom/parser"
	"github.com/pkg/errors"
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

	definitions, err := parser.Parse(xml, comments)
	if err != nil {
		return err
	}

	// start with enums

	for _, d := range definitions {
		if t, ok := d.(defs.Interface); ok {
			if t.Name() != "Node" {
				continue
			}
			// imps, e := t.ImplementedBy()
			// if e != nil {
			// 	return e
			// }
			// if len(imps) > 0 {
			// 	fmt.Println("implements=%s", t.Name())
			// }
			str, err := t.Generate()
			if err != nil {
				return errors.Wrap(err, "error generating")
			}
			fmt.Println(str)
		}
		// children, err := d.Children()
		// if err != nil {
		// 	return err
		// }
		// _ = children
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
