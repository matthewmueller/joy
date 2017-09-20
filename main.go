package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/matthewmueller/golly/js"
)

// Test struct
type Test struct {
	Name      string `json:"name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
}

func main() {
	log.SetHandler(text.New(os.Stderr))

	file := os.Args[1]
	gosrc, err := ioutil.ReadFile(file)
	if err != nil {
		log.WithError(err).Fatalf("couldn't read file")
	}

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, file, gosrc, 0)
	if err != nil {
		panic(err)
	}

	

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		fmt.Println(reflect.TypeOf(n))
		// var s string
		// switch x := n.(type) {
		// case *ast.BasicLit:
		// 	s = x.Value
		// case *ast.Ident:
		// 	s = x.Name
		// default:
		// 	log.Infof("missed %s", reflect.TypeOf(n))
		// }
		// if s != "" {
		// 	fmt.Printf("%s:\t%s\n", fset.Position(n.Pos()), s)
		// }
		return true
	})

	// ast.Inspect(go)
	// "type": "Program",
	// "start": 0,
	// "end": 47,
	// "body": [
	//   {
	//     "type": "EmptyStatement",
	//     "start": 0,
	//     "end": 1
	//   },

	program := js.CreateProgram(
		js.CreateEmptyStatement(),
		js.CreateExpressionStatement(
			js.CreateCallExpression(
				js.CreateFunctionExpression(nil, []js.IPattern{},
					js.CreateFunctionBody(
						js.CreateExpressionStatement(
							js.CreateCallExpression(
								js.CreateMemberExpression(
									js.CreateIdentifier("console"),
									js.CreateIdentifier("log"),
									false,
								),
								[]js.IExpression{
									js.CreateString("hi world!"),
								},
							),
						),
					),
				),
				[]js.IExpression{},
			),
		),
	)

	// program := js.Program{
	// 	Type: "Program",
	// 	Body: []interface{}{
	// 		js.EmptyStatement{
	// 			Type: "EmptyStatement",
	// 		},
	// 	},
	// }

	jssrc, err := js.Generate(program)
	if err != nil {
		log.WithError(err).Fatalf("syntax error")
	}
	fmt.Println(jssrc)

	// ast, err := json.MarshalIndent(program, "", "  ")
	// if err != nil {
	// 	log.WithError(err).Fatalf("couldn't marshal json")
	// }

	// fmt.Println(string(ast))
}
