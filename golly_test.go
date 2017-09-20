package golly_test

import (
	"fmt"
	"io/ioutil"
	"path"
	"testing"

	"github.com/matthewmueller/golly"
	"github.com/stretchr/testify/assert"
)

func TestCompiler(t *testing.T) {
	dirs, err := ioutil.ReadDir("./testfiles")
	if err != nil {
		t.Fatal(err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		// subtests
		t.Run(dir.Name(), func(t *testing.T) {
			input := path.Join("./testfiles", dir.Name(), "input.go")
			output := path.Join("./testfiles", dir.Name(), "output.js.txt")

			// read the output
			js, e := ioutil.ReadFile(output)
			if e != nil {
				t.Fatal(e)
			}

			// compile the file
			out, e := golly.CompileFile(input)
			if e != nil {
				golly.PrintAST(input)
				t.Fatal(e)
			}

			if string(js) != out {
				golly.PrintAST(input)
				fmt.Println(out)
			}

			assert.Equal(t, string(js), out)
		})
	}
}

// func TestConsole(t *testing.T) {
// 	// Create the AST by parsing src.
// 	fset := token.NewFileSet() // positions are relative to fset
// 	f, err := parser.ParseFile(fset, "main.go", nil, 0)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Inspect the AST and print all identifiers and literals.
// 	ast.Inspect(f, func(n ast.Node) bool {
// 		var s string
// 		switch x := n.(type) {
// 		case *ast.BasicLit:
// 			s = x.Value
// 		case *ast.Ident:
// 			s = x.Name
// 		}
// 		if s != "" {
// 			fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
// 		}
// 		return true
// 	})
// }
