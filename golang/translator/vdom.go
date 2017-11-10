package translator

import (
	"go/ast"
	"strings"

	"github.com/apex/log"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/defs"
	"github.com/matthewmueller/golly/golang/util"
	"github.com/matthewmueller/golly/jsast"
)

func (tr *Translator) maybeVDOMLit(d def.Definition, n *ast.CompositeLit) (jsast.IExpression, error) {
	def, err := tr.index.DefinitionOf(d.Path(), n)
	if err != nil {
		return nil, err
	}

	stct, ok := def.(defs.Structer)
	if !ok {
		return nil, nil
	}

	jsxPath, err := util.JSXSourcePath()
	if err != nil {
		return nil, err
	}

	ifaces, err := stct.Implements()
	if err != nil {
		return nil, err
	}

	isNode := false
	for _, iface := range ifaces {
		if iface.Path() == jsxPath && iface.Name() == "Node" {
			isNode = true
			break
		}
	}
	if !isNode {
		return nil, nil
	}

	pragma, err := tr.index.JSXPragma()
	if err != nil {
		return nil, err
	}

	kvs, e := tr.getKeyValues(d, nil, n)
	if e != nil {
		return nil, e
	}

	// handle text nodes differently
	if stct.OriginalName() == "Text" {
		for key, val := range kvs {
			if strings.ToLower(key) == "value" {
				return tr.expression(d, nil, val)
			}
		}
		return jsast.CreateString(""), nil
	}

	// look for props if we have them
	var props []jsast.Property
	for key, val := range kvs {
		if strings.ToLower(key) == "props" {
			c, ok := val.(*ast.CompositeLit)
			if !ok {
				// TODO: should we error here? or just ignore?
				continue
				// return nil, fmt.Errorf("translator/jsNewFunction: expected props to be a compositLiteral, but got %T", val)
			}

			for i, elt := range c.Elts {
				p, e := tr.property(d, nil, c, i, elt)
				if e != nil {
					return nil, e
				}
				props = append(props, p)
			}
		}
	}

	fn, e := tr.expression(d, nil, n.Type)
	if e != nil {
		return nil, e
	}

	return jsast.CreateCallExpression(
		jsast.CreateIdentifier(pragma),
		[]jsast.IExpression{
			fn,
			jsast.CreateObjectExpression(props),
		},
	), nil
}

func (tr *Translator) maybeVDOMReturn(d def.Definition, n *ast.ReturnStmt) (j jsast.IStatement, err error) {
	fn, ok := d.(defs.Functioner)
	if !ok {
		return nil, nil
	}

	idx, err := findComponentIndex(fn)
	if err != nil {
		return nil, err
	} else if idx < 0 {
		return nil, nil
	}

	var args []jsast.IExpression
	for i, arg := range n.Results {
		a, e := tr.expression(d, nil, arg)
		if e != nil {
			return nil, e
		}

		// wrap the return in a sequence expression
		if i == idx {
			id, e := util.GetIdentifier(arg)
			if e != nil {
				return nil, e
			}
			s, e := util.ExprToString(id)
			if e != nil {
				return nil, e
			}
			log.Infof("s=%s", s)

			a = jsast.CreateSequenceExpression(
				a,
				jsast.CreateRaw("preact.h("+s+", "+s+".props)"),
			)
		}

		args = append(args, a)
	}

	// return an array
	if len(n.Results) > 1 {
		return jsast.CreateReturnStatement(jsast.CreateArrayExpression(args...)), nil
	}

	// return the value by itself
	return jsast.CreateReturnStatement(args[0]), nil
}

func findComponentIndex(fn defs.Functioner) (int, error) {
	jsxPath, err := util.JSXSourcePath()
	if err != nil {
		return -1, err
	}

	for i, result := range fn.Results() {
		def := result.Definition()
		if def == nil {
			continue
		}

		switch t := def.(type) {
		case defs.Interfacer:
			if t.Path() == jsxPath && t.Name() == "Node" {
				return i, nil
			}
		case defs.Structer:
			imps, err := t.Implements()
			if err != nil {
				return -1, err
			}
			for _, imp := range imps {
				if imp.Path() == jsxPath && imp.Name() == "Node" {
					return i, nil
				}
			}
		}
	}

	return -1, nil
}
