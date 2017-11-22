//go:generate go run main.go

package main

import (
	"github.com/apex/log"
	"github.com/matthewmueller/golly/internal/dom/curl"
	"github.com/matthewmueller/golly/internal/dom/def"
	"github.com/matthewmueller/golly/internal/dom/graph"
	"github.com/matthewmueller/golly/internal/dom/parser"
	"github.com/pkg/errors"
)

var browserAPIURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/browser.webidl.xml"
var browserDocsURL = "https://rawgit.com/Microsoft/TSJS-lib-generator/master/inputfiles/comments.json"

//
// var packages = map[string][]string {
// 	"dom":
// }

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
	visited := map[string]bool{}
	queue := []def.Definition{}
	// start with enums

	queue = append(queue,
		definitions["Attr"],
		definitions["CharacterData"],
		definitions["ChildNode"],
		definitions["Comment"],
		definitions["CustomEvent"],
		definitions["Document"],
		definitions["DocumentFragment"],
		definitions["DocumentType"],
		definitions["DOMError"],
		definitions["DOMException"],
		definitions["DOMImplementation"],
		definitions["DOMSettableTokenList"],
		definitions["DOMStringList"],
		definitions["DOMTokenList"],
		definitions["Element"],
		definitions["Event"],
		definitions["EventTarget"],
		definitions["HTMLCollection"],
		definitions["MutationObserver"],
		definitions["MutationRecord"],
		definitions["NamedNodeMap"],
		definitions["Node"],
		definitions["NodeFilter"],
		definitions["NodeIterator"],
		definitions["NodeList"],
		definitions["ProcessingInstruction"],
		definitions["Selection"],
		definitions["Range"],
		definitions["Text"],
		definitions["TimeRanges"],
		definitions["TreeWalker"],
		definitions["URL"],
		definitions["Window"],
		definitions["Worker"],
		definitions["XMLDocument"],
	)

	for _, def := range queue {
		visited[def.ID()] = true
	}

	for len(queue) > 0 {
		// dequeue
		d := queue[0]
		queue = queue[1:]

		deps, err := d.Dependencies()
		if err != nil {
			return errors.Wrap(err, "deps")
		}
		deps = unique(deps)

		// add the dependencies to the graph
		for _, dep := range deps {
			g.Edge(d, dep)

			// queue up any dependency that hasn't
			// already been visited
			if !visited[dep.ID()] {
				queue = append(queue, dep)
				visited[d.ID()] = true
			}
		}
	}

	// def := definitions["Element"]
	// deps, err := def.Dependencies()
	// if err != nil {
	// 	return errors.Wrap(err, "error getting deps")
	// }
	// for _, dep := range deps {
	// 	log.Infof(dep.Name())
	// }

	cliques := g.Cliques()
	for _, clique := range cliques {
		log.Infof("clique len(%d)", len(clique))
		for _, def := range clique {
			log.Infof("def=%s", def.ID())
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
