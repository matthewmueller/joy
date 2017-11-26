package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/alonsovidales/go_graph"

	"github.com/apex/log"
)

func main() {
	cwd, e := os.Getwd()
	if e != nil {
		log.WithError(e).Fatalf("fatal")
	}

	base, e := filepath.Rel(path.Join(os.Getenv("GOPATH"), "src"), cwd)
	if e != nil {
		log.WithError(e).Fatalf("rel")
	}

	depmap := map[string][]string{}
	intmap := map[string]uint64{}
	strmap := map[uint64]string{}
	var index uint64 = 1

	e = filepath.Walk("./internal/gendom/dom", func(p string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		buf, e := ioutil.ReadFile(p)
		if e != nil {
			return e
		}

		fset := token.NewFileSet()
		file, e := parser.ParseFile(fset, p, buf, parser.ImportsOnly)
		if e != nil {
			return e
		}

		parent := path.Join(base, path.Dir(p))
		if depmap[parent] == nil {
			depmap[parent] = []string{}
			intmap[parent] = index
			strmap[index] = parent
			index++
		}

		for _, imp := range file.Imports {
			s, e := strconv.Unquote(imp.Path.Value)
			if e != nil {
				return e
			}

			// only graph internal gendom deps
			if !strings.Contains(s, "/internal/gendom/dom") {
				continue
			}

			depmap[parent] = append(depmap[parent], s)
		}
		return nil
	})
	if e != nil {
		log.WithError(e).Errorf("error walking")
	}

	var edges []graphs.Edge
	for parent, deps := range depmap {
		// log.Infof("- %s", parent)
		for _, dep := range deps {
			if intmap[parent] == 0 {
				panic("parent shouldn't be 0: " + parent)
			}
			if intmap[dep] == 0 {
				panic("dep shouldn't be 0: " + dep)
			}

			edges = append(edges, graphs.Edge{
				From: intmap[parent],
				To:   intmap[dep],
			})
		}
	}

	g := graphs.GetGraph(edges, false)

	components, groups := g.StronglyConnectedComponents()
	for _, group := range groups {
		fmt.Printf("group size: %d\n", len(group))
		for key := range group {
			fmt.Printf("    %s\n", strmap[key])
		}
	}
	_ = components
	// components := scc.StronglyConnectedComponents(graph)
	// for _, component := range components {
	// 	// log.Infof("")
	// }

	// strongly_connected_components.Graph{}
}
