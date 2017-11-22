//go:generate go run main.go

package main

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/internal/dom/curl"
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/graph"
	"github.com/matthewmueller/golly/internal/dom/parser"
	"github.com/matthewmueller/golly/internal/gen"
	"github.com/pkg/errors"
)

var browserAPIURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/browser.webidl.xml"
var browserDocsURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/comments.json"

// manually selected package names for the cliques
var cliqueNames = map[string]string{
	"IntersectionObserver": "intersectionobserver",
	"MutationObserver":     "mutationobserver",
	"MimeType":             "mimetype",
	"Window":               "window",
	"MediaQueryList":       "mediaquery",
	"IDBDatabase":          "idb",
	"VideoTrack":           "avtrack",
	"TextTrack":            "texttrack",
	"MSHTMLWebViewElement": "mswebviewelement",
	"SVGUseElement":        "svguseelement",
}

func generate(dir string) error {
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

	g := graph.New()
	for _, parent := range definitions {
		deps, err := parent.Dependencies()
		if err != nil {
			return err
		}
		for _, child := range deps {
			g.Edge(parent, child)
		}
	}

	// generate the packages
	cliques := g.Cliques()
	for _, clique := range cliques {
		if len(clique) == 1 {
			name := clique[0].ID()
			pkgname := gen.Lowercase(name)
			def := definitions[name]
			def.SetPackage(pkgname)
			def.SetFile(pkgname)
			continue
		}

		pkgname := ""
		for _, def := range clique {
			if cliqueNames[def.ID()] != "" {
				pkgname = cliqueNames[def.ID()]
				break
			}
		}
		if pkgname == "" {
			return errors.New("clique name is not defined")
		}

		for _, def := range clique {
			name := def.ID()
			filename := gen.Lowercase(name)
			def := definitions[name]
			def.SetPackage(pkgname)
			def.SetFile(filename)
		}
	}

	for id, def := range definitions {
		// if id != "window" {
		// 	continue
		// }

		log.Infof("package %s", id)

		code, err := def.Generate()
		if err != nil {
			return errors.Wrapf(err, "error generating %s", id)
		}

		if err := os.MkdirAll(path.Join(dir, def.GetPackage()), 0755); err != nil {
			return errors.Wrapf(err, "error mkdir")
		}

		if err := ioutil.WriteFile(def.GetFile()+".go", []byte(code), 0644); err != nil {
			return errors.Wrapf(err, "error writefile")
		}
	}

	return nil
}

func unique(defs []def.Definition) []def.Definition {
	u := make([]def.Definition, 0, len(defs))
	m := make(map[string]bool)

	for _, def := range defs {
		if _, ok := m[def.ID()]; !ok {
			m[def.ID()] = true
			u = append(u, def)
		}
	}

	return u
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	pwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatalf("error getting cwd")
	}

	if e := generate(path.Join(pwd, "dom2")); e != nil {
		log.WithError(e).Fatalf("error generating")
	}
}
