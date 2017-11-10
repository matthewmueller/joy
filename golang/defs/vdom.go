package defs

import (
	"go/ast"
	"strconv"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/util"
	"github.com/pkg/errors"
)

func jsxUse(ctx *context, n *ast.CallExpr) error {
	if len(n.Args) < 2 {
		return errors.New("process/jsxUse: expected atleast 2 arguments")
	}

	p, err := util.ExprToString(n.Args[0])
	if err != nil {
		return errors.Wrap(err, "process/jsxUse: error getting pragma")
	}
	pragma, err := strconv.Unquote(p)
	if err != nil {
		return errors.Wrap(err, "process/jsxUse: couldn't unquote pragma")
	}

	f, err := util.ExprToString(n.Args[1])
	if err != nil {
		return errors.Wrap(err, "process/jsxUse: error getting filepath")
	}
	filepath, err := strconv.Unquote(f)
	if err != nil {
		return errors.Wrap(err, "process/jsxUse: couldn't unquote filepath")
	}

	// add to the index
	ctx.idx.SetJSXPragma(pragma)

	// add the file dependency
	def, e := File(ctx.d.Path(), filepath)
	if e != nil {
		return e
	}
	ctx.idx.AddDefinition(def)
	log.Debugf("%s -> %s", ctx.d.ID(), def.ID())
	ctx.state.deps = append(ctx.state.deps, def)

	return nil
}

func maybeVDOMCompositLit(ctx *context, n *ast.CompositeLit) error {
	def, err := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if err != nil {
		return err
	}

	stct, ok := def.(Structer)
	if !ok {
		return nil
	}

	jsxPath, err := util.JSXSourcePath()
	if err != nil {
		return err
	}

	ifaces, err := stct.Implements()
	if err != nil {
		return err
	}

	isNode := false
	for _, iface := range ifaces {
		if iface.Path() == jsxPath && iface.Name() == "Node" {
			isNode = true
			break
		}
	}
	if !isNode {
		return nil
	}

	for _, method := range stct.Methods() {
		ctx.state.deps = append(
			ctx.state.deps,
			method,
		)
	}

	// pragma, err := ctx.idx.JSXPragma()
	// if err != nil {
	// 	return err
	// }

	// log.Infof("name=%s pragma=%s", stct.OriginalName(), pragma)

	return nil
}
