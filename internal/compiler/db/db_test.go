package db_test

// import (
// 	"go/ast"
// 	"io/ioutil"
// 	"os"
// 	"path"
// 	"path/filepath"
// 	"strings"
// 	"testing"

// 	"github.com/apex/log"
// 	"github.com/apex/log/handlers/text"
// 	"github.com/matthewmueller/joy/internal/compiler/db"
// 	"github.com/matthewmueller/joy/internal/compiler/def/iface"
// 	"github.com/matthewmueller/joy/internal/compiler/loader"
// )

// func Test(t *testing.T) {
// 	// ctx := context.Background()
// 	log.SetHandler(text.New(os.Stderr))

// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	dirs, err := ioutil.ReadDir("../../testdata")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	for _, dir := range dirs {
// 		if !dir.IsDir() {
// 			continue
// 		}

// 		// subtests
// 		t.Run(dir.Name(), func(t *testing.T) {
// 			if strings.HasPrefix(dir.Name(), "_") {
// 				log.Warnf("Skipping '%s'", dir.Name())
// 				return
// 			}

// 			testdir := path.Join(cwd, "../../testdata", dir.Name())
// 			var inputs []string

// 			// support single and multipage tests
// 			inps, e := filepath.Glob(path.Join(testdir, "input.go"))
// 			if e != nil {
// 				t.Fatal(e)
// 			}
// 			inputs = append(inputs, inps...)

// 			inps, e = filepath.Glob(path.Join(testdir, "*", "input.go"))
// 			if e != nil {
// 				t.Fatal(e)
// 			}
// 			inputs = append(inputs, inps...)

// 			var pages []string
// 			for _, input := range inputs {
// 				pages = append(pages, path.Dir(input))
// 				log.Infof("page %s", path.Dir(input))
// 			}
// 			log.Infof("inputs %+v", inputs)

// 			program, err := loader.Load(pages...)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			database, err := db.New(program)
// 			if err != nil {
// 				t.Fatal(err)
// 			}

// 			for _, info := range program.InitialPackages() {
// 				for _, file := range info.Files {
// 					ast.Inspect(file, func(node ast.Node) bool {
// 						switch n := node.(type) {
// 						case *ast.Ident:
// 							d, e := database.DefinitionOf(info, n)
// 							if e != nil {
// 								t.Fatal(e)
// 							}
// 							if d != nil {
// 								log.Infof("found %s:", d.ID())
// 							} else {
// 								// log.Infof("not found %s .%s %T", info.Pkg.Path(), n.Sel.Name, n)
// 								return false
// 							}

// 							// if m, ok := d.(def.Method); ok {
// 							// 	recv := m.Recv()
// 							// 	if recv == nil {
// 							// 		log.Infof("recv not found in %s: %s", d.ID(), m.Recv())
// 							// 	} else {
// 							// 		log.Infof("recv found! %s", recv.ID())
// 							// 	}
// 							// }

// 							if i, ok := d.(iface.Interface); ok {
// 								imps := i.ImplementedBy("String")
// 								for _, imp := range imps {
// 									log.Infof("%s implemented by %s", i.ID(), imp.ID())
// 								}
// 							}
// 						}
// 						return true
// 					})
// 				}
// 			}
// 		})
// 	}
// }
