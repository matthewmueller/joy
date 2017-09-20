package main

import (
	"fmt"
	"io/ioutil"
	"os"

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

	gosrc, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.WithError(err).Fatalf("couldn't read from stdin")
	}
	_ = gosrc

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
