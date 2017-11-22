//go:generate go run main.go

package main

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/dom/curl"
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/graph"
	"github.com/matthewmueller/golly/internal/dom/parser"
)

var browserAPIURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/browser.webidl.xml"
var browserDocsURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/comments.json"

// manually selected package names for the cliques
var cliques = map[string]string{
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

	cliques := g.Cliques()
	for _, clique := range cliques {
		if len(clique) > 1 {
			log.Infof("clique len(%d)", len(clique))
			for _, def := range clique {
				log.Infof("def=%s", def.ID())
			}
		}
	}

	// for _, d := range definitions {
	// 	if t.
	// 	if t, ok := d.(defs.Interface); ok {
	// 		if t.Name() != "Window" {
	// 			continue
	// 		}

	// 		queue = append(queue, t)
	// 		visited[t.ID()] = true

	// 		// imps, e := t.ImplementedBy()
	// 		// if e != nil {
	// 		// 	return e
	// 		// }
	// 		// if len(imps) > 0 {
	// 		// 	fmt.Println("implements=%s", t.Name())
	// 		// }
	// 		str, err := t.Generate()
	// 		if err != nil {
	// 			return errors.Wrap(err, "error generating")
	// 		}
	// 		fmt.Println(str)
	// 	}
	// 	// children, err := d.Children()
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	// _ = children
	// 	// log.Infof("got children %d", )
	// }

	// attr := idx["Attr"]
	// log.Infof("attr %+v", attr.Children())

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
	if e := generate(); e != nil {
		log.WithError(e).Fatalf("error generating")
	}
}
