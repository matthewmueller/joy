package translator

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/matthewmueller/joy/internal/compiler/def"
	"github.com/matthewmueller/joy/internal/compiler/defs"
	"github.com/matthewmueller/joy/internal/compiler/index"
	"github.com/matthewmueller/joy/internal/compiler/scope"
	"github.com/matthewmueller/joy/internal/compiler/util"
	"github.com/matthewmueller/joy/internal/jsast"
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

	vdomPath, err := util.VDOMSourcePath()
	if err != nil {
		return nil, err
	}

	ifaces, err := stct.Implements()
	if err != nil {
		return nil, err
	}

	isNode := false
	for _, iface := range ifaces {
		if iface.Path() == vdomPath && iface.Name() == "Component" {
			isNode = true
			break
		}
	}
	if !isNode {
		return nil, nil
	}

	var propsName string
	fields := stct.Fields()
	for _, field := range fields {
		if field.Name() == "props" {
			propsStruct, err := tr.index.DefinitionOf(d.Path(), field.Type())
			if err != nil {
				return nil, err
			} else if propsStruct == nil || propsStruct.Kind() != "STRUCT" {
				return nil, fmt.Errorf("maybeVDOMLit: expected the props field to point to a struct")
			}
			propsName = propsStruct.Name()
		}
	}
	if propsName == "" {
		return nil, fmt.Errorf(`vdom: "%s" struct must have a field named "props"`, stct.Name())
	}

	kvs, e := tr.getKeyValues(d, nil, n)
	if e != nil {
		return nil, e
	}

	// look for props if we have them
	var props jsast.IExpression
	for key, val := range kvs {
		if strings.ToLower(key) == "props" {
			c := findCompositLit(val)
			if c == nil {
				continue
			}

			cl, err := tr.compositeLiteral(d, nil, c)
			if err != nil {
				return nil, err
			}
			props = cl
		}
	}
	if props == nil {
		props = jsast.CreateNewExpression(
			jsast.CreateIdentifier(propsName),
			[]jsast.IExpression{jsast.CreateObjectExpression(nil)},
		)
	}

	fn, e := tr.expression(d, nil, n.Type)
	if e != nil {
		return nil, e
	}

	pragma, err := resolveVDOMPragma(tr.index)
	if err != nil {
		return nil, err
	}

	return jsast.CreateCallExpression(
		pragma,
		[]jsast.IExpression{
			fn,
			props,
		},
	), nil
}

func (tr *Translator) vdomPragma() (j jsast.IExpression, err error) {
	return resolveVDOMPragma(tr.index)
}

func (tr *Translator) vdomFile() (j jsast.IExpression, err error) {
	return resolveVDOMFile(tr.index)
}

func (tr *Translator) vdomComponent(d def.Definition, sp *scope.Scope, n *ast.SelectorExpr) (j jsast.IExpression, err error) {
	pkg, err := resolveVDOMFile(tr.index)
	if err != nil {
		return nil, err
	}

	return jsast.CreateMemberExpression(
		pkg,
		jsast.CreateIdentifier("Component"),
		false,
	), nil
}

func resolveVDOMFile(idx *index.Index) (j jsast.MemberExpression, e error) {
	file, err := idx.VDOMFile()
	if err != nil {
		return j, err
	}

	return jsast.CreateMemberExpression(
		jsast.CreateIdentifier("pkg"),
		jsast.CreateString(file.Path()),
		true,
	), nil
}

func resolveVDOMPragma(idx *index.Index) (j jsast.MemberExpression, e error) {
	pragma, err := idx.VDOMPragma()
	if err != nil {
		return j, err
	}

	parts := strings.Split(pragma, ".")
	remaining := strings.Join(parts[1:], ".")

	file, err := idx.VDOMFile()
	if err != nil {
		return j, err
	}

	return jsast.CreateMemberExpression(
		jsast.CreateMemberExpression(
			jsast.CreateIdentifier("pkg"),
			jsast.CreateString(file.Path()),
			true,
		),
		jsast.CreateIdentifier(remaining),
		false,
	), nil
}

func findComponentIndex(fn defs.Functioner) (int, error) {
	vdomPath, err := util.VDOMSourcePath()
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
			if t.Path() == vdomPath && t.Name() == "Component" {
				return i, nil
			}
		case defs.Structer:
			imps, err := t.Implements()
			if err != nil {
				return -1, err
			}
			for _, imp := range imps {
				if imp.Path() == vdomPath && imp.Name() == "Component" {
					return i, nil
				}
			}
		}
	}

	return -1, nil
}

func findCompositLit(n ast.Expr) *ast.CompositeLit {
	switch t := n.(type) {
	case *ast.CompositeLit:
		return t
	case *ast.UnaryExpr:
		return findCompositLit(t.X)
	}
	return nil
}
