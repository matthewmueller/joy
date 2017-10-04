package disassembler

import (
	"github.com/apex/log"
	"github.com/dop251/goja"
	"github.com/matthewmueller/golly/bindata"
	"github.com/matthewmueller/golly/jsast"
	"github.com/pkg/errors"
)

// Acorn struct
type Acorn struct {
	VM   *goja.Runtime
	Fn   goja.Callable
	This goja.Value
}

// Disassemble a file into AST nodes
func Disassemble(src string) (program jsast.Program, err error) {
	acorn, err := bindata.Asset("bindata/acorn.js")
	if err != nil {
		return program, err
	}

	vm := goja.New()
	prg, err := goja.Compile("acorn.js", string(acorn), false)
	if err != nil {
		return program, errors.Wrap(err, "couldn't compile acorn.js")
	}

	_, err = vm.RunProgram(prg)
	if err != nil {
		return program, errors.Wrap(err, "couldn't run acorn program")
	}
	log.Infof("%+v", vm.Get("acorn"))

	value := vm.Get("acorn")
	exported := value.ToObject(vm)

	log.Infof("exported %+v", exported.Get("parse"))

	// fn, ok := goja.AssertFunction(vm.Get("global"))
	// if !ok {
	// 	return program, errors.New("couldn't get the acorn function")
	// }

	// log.Infof("got function %+v", fn)

	return program, nil
}
