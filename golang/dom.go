package golang

import (
	"github.com/apex/log"
	"github.com/kr/pretty"
	"github.com/matthewmueller/golly/js"
)

// The DOM's universe
// var Universe = map[string]func()

// DOM mappings
var DOM = map[string]func(interface{}) (js.IExpression, error){
	"dom.Document": func(v interface{}) (js.IExpression, error) {
		pretty.Println(v)
		log.Infof("here?!?!?")
		return nil, nil
	},
}
