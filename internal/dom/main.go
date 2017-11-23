//go:generate go run main.go

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

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
	"WebKitFileSystem":     "webkitfilesytem",
	"AudioNode":            "audionode",
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

	index, err := parser.Parse(xml, comments)
	if err != nil {
		return err
	}

	g := graph.New()
	for _, parent := range index {
		g.Node(parent)

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
			def := index[name]
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
			var ids []string
			for _, def := range clique {
				ids = append(ids, def.ID())
			}
			return fmt.Errorf("group name not defined for this clique: %s", strings.Join(ids, ", "))
		}

		for _, def := range clique {
			name := def.ID()
			filename := gen.Lowercase(name)
			def := index[name]
			def.SetPackage(pkgname)
			def.SetFile(filename)
		}
	}

	// sort so goimports finds dependencies
	var defs []def.Definition
	sorted, _ := g.Toposort()
	for _, node := range sorted {
		def := index[node.ID()]

		// only use these as these are our top-level files
		switch def.Kind() {
		case "ENUM", "DICTIONARY", "INTERFACE":
			defs = append(defs, def)
		}
	}

	l := len(defs)
	for i, def := range defs {
		code, err := def.Generate()
		if err != nil {
			return errors.Wrapf(err, "error generating %s", def.ID())
		}

		formatted, err := gen.Format(code)
		if err != nil {
			return errors.Wrapf(err, "error formatting %s received\n%s", def.ID(), code)
		}

		pkgpath := path.Join(dir, def.GetPackage())
		if err := os.MkdirAll(pkgpath, 0755); err != nil {
			return errors.Wrapf(err, "error mkdir")
		}

		if err := ioutil.WriteFile(path.Join(pkgpath, def.GetFile()+".go"), []byte(formatted), 0644); err != nil {
			return errors.Wrapf(err, "error writefile")
		}

		log.Infof("generated %s (%d/%d)", def.ID(), i, l)
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

	if err := os.RemoveAll(path.Join(pwd, "dom2")); err != nil {
		log.WithError(err).Fatalf("removing dom")
	}

	if e := generate(path.Join(pwd, "dom2")); e != nil {
		log.WithError(e).Fatalf("error generating")
	}
}
