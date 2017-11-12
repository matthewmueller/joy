package defs

import (
	"go/ast"
	"strconv"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/util"
	"github.com/pkg/errors"
)

// https://reactjs.org/docs/react-component.html
var componentAPI = map[string]string{
	"ComponentWillMount":        "componentWillMount",
	"Render":                    "render",
	"ComponentDidMount":         "componentDidMount",
	"ComponentWillReceiveProps": "componentWillReceiveProps",
	"ShouldComponentUpdate":     "shouldComponentUpdate",
	"ComponentWillUpdate":       "componentWillUpdate",
	"ComponentDidUpdate":        "componentDidUpdate",
	"ComponentWillUnmount":      "componentWillUnmount",
	"ComponentDidCatch":         "componentDidCatch",
}

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

	// add the file dependency
	file, e := File(ctx.d.Path(), filepath, true)
	if e != nil {
		return e
	}

	ctx.idx.SetJSXFile(file)

	// add to the index
	ctx.idx.SetJSXPragma(pragma)

	return nil
}

func maybeVDOMCompositLit(ctx *context, n *ast.CompositeLit) error {
	def, err := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if err != nil {
		return err
	} else if def == nil {
		return nil
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

	// add the public react API
	for _, method := range stct.Methods() {
		_, isset := componentAPI[method.OriginalName()]
		if !isset {
			continue
		}
		ctx.state.deps = append(
			ctx.state.deps,
			method,
		)
	}

	file, err := ctx.idx.JSXFile()
	if err != nil {
		return err
	}

	pragma, err := ctx.idx.JSXPragma()
	if err != nil {
		return err
	}

	// get the first part of the pragma
	// e.g.
	//   preact.h => preact
	//   React.createElement => React
	ids := strings.SplitN(pragma, ".", 2)

	ctx.idx.AddDefinition(file)
	log.Debugf("%s -> %s", ctx.d.ID(), file.ID())
	ctx.state.deps = append(ctx.state.deps, file)
	ctx.state.imports[ids[0]] = file.Path()

	return nil
}

func maybeVDOMFuncDecl(ctx *context, n *ast.FuncDecl) error {
	def, err := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if err != nil {
		return err
	} else if def == nil {
		return nil
	}

	method, ok := def.(Methoder)
	if !ok {
		return nil
	}

	recv := method.Recv()
	if recv == nil {
		return nil
	}

	stct, ok := recv.(Structer)
	if !ok {
		return errors.New("maybeVDOMFuncDecl: expected receiver to be a struct")
	}

	jsxPath, err := util.JSXSourcePath()
	if err != nil {
		return err
	}

	defs, err := stct.Implements()
	if err != nil {
		return err
	}

	isComponent := false
	for _, def := range defs {
		if def.OriginalName() == "Node" && jsxPath == def.Path() {
			isComponent = true
		}
	}
	if !isComponent {
		return nil
	}

	// rename according to the componentAPI
	alias, isset := componentAPI[def.OriginalName()]
	if !isset {
		return nil
	}
	ctx.state.rename = alias

	return nil
}
