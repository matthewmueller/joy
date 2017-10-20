package defs

import (
	"fmt"
	"go/ast"
	"strconv"

	"github.com/apex/log"
	"github.com/fatih/structtag"
	"github.com/matthewmueller/golly/golang/def"
	"github.com/matthewmueller/golly/golang/index"
	"github.com/matthewmueller/golly/golang/util"
)

type context struct {
	idx   *index.Index
	d     def.Definition
	state *state
}

type state struct {
	imports map[string]string
	tag     *structtag.Tag
	deps    []def.Definition
	rewrite *rewrite
	fields  []def.Field
	async   bool
	omit    bool
}

// private processor for each of the definition types
func process(idx *index.Index, d def.Definition, n ast.Node) (st *state, err error) {
	st = &state{
		imports: map[string]string{},
	}

	ctx := &context{
		state: st,
		idx:   idx,
		d:     d,
	}

	ast.Inspect(n, func(n ast.Node) bool {
		switch t := n.(type) {
		case *ast.FuncDecl:
			err = funcDecl(ctx, t)
		case *ast.GenDecl:
			err = genDecl(ctx, t)
		case *ast.CallExpr:
			err = callExpr(ctx, t)
		case *ast.SelectorExpr:
			err = selectorExpr(ctx, t)
		case *ast.ChanType:
			err = chanType(ctx, t)
		case *ast.Ident:
			err = ident(ctx, t)
		}
		return err == nil
	})
	if err != nil {
		return nil, err
	}

	// unique dependencies
	st.deps = unique(st.deps)

	return st, nil
}

func funcDecl(ctx *context, n *ast.FuncDecl) error {
	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return e
	}
	ctx.state.tag = tag
	return nil
}

func genDecl(ctx *context, n *ast.GenDecl) error {
	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return e
	}
	ctx.state.tag = tag
	return nil
}

func structType(ctx *context, n *ast.StructType) error {
	for _, f := range n.Fields.List {
		fld := &field{kind: f.Type}

		// add a tag if we have one
		if f.Tag != nil {
			v, e := strconv.Unquote(f.Tag.Value)
			if e != nil {
				return e
			}
			tag, e := util.JSTagFromString(v)
			if e != nil {
				return e
			}
			fld.tag = tag
		}

		// handle User{*dep.Settings} w/o name
		if len(f.Names) == 0 {
			id, e := util.GetIdentifier(f.Type)
			if e != nil {
				return e
			}
			fld.name = id.Name
			ctx.state.fields = append(ctx.state.fields, fld)
			continue
		}

		for _, id := range f.Names {
			ctx.state.fields = append(ctx.state.fields, &field{
				name: id.Name,
				kind: fld.kind,
				tag:  fld.tag,
			})
		}
	}

	return nil
}

func callExpr(ctx *context, n *ast.CallExpr) error {
	cx, e := util.ExprToString(n.Fun)
	if e != nil {
		return e
	}

	switch cx {
	case "js.RawFile":
		return jsFile(ctx, n)
	case "js.Rewrite":
		return jsRewrite(ctx, n)
	}

	// TODO: handle interfaces
	def, e := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if e != nil {
		return e
	}

	switch t := def.(type) {
	case Interfacer:
		log.Infof("interface! cx=%s def=%s", cx, t.ID())
	}

	return nil
}

func chanType(ctx *context, n *ast.ChanType) error {
	deps, e := ctx.idx.Runtime("Deferred", "Channel", "Send", "Recv")
	if e != nil {
		return e
	}

	if len(deps) > 0 {
		ctx.state.deps = append(ctx.state.deps, deps...)
		ctx.state.imports["runtime"] = deps[0].Path()
	}

	// TODO: more specific later
	ctx.state.async = true
	return nil
}

func selectorExpr(ctx *context, n *ast.SelectorExpr) error {
	def, e := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if e != nil {
		return e
	}

	// add dependencies
	if def != nil {
		ctx.state.deps = append(ctx.state.deps, def)
	}

	// handle async
	if def != nil && !ctx.state.async && ctx.d.ID() != def.ID() {
		async, e := isAsync(def)
		if e != nil {
			return e
		}
		ctx.state.async = async
	}

	return nil
}

func ident(ctx *context, n *ast.Ident) error {
	def, e := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if e != nil {
		return e
	}

	// add dependencies
	if def != nil {
		ctx.state.deps = append(ctx.state.deps, def)
	}

	// handle async
	if def != nil && !ctx.state.async && ctx.d.ID() != def.ID() {
		async, e := isAsync(def)
		if e != nil {
			return e
		}
		ctx.state.async = async
	}

	return nil
}

func jsFile(ctx *context, n *ast.CallExpr) error {
	if len(n.Args) == 0 {
		return nil
	}

	lit, ok := n.Args[0].(*ast.BasicLit)
	if !ok {
		return fmt.Errorf("fn process: expected rawfile to have basiclit argument, but got %T", n.Args[0])
	}

	path, e := strconv.Unquote(lit.Value)
	if e != nil {
		return e
	}

	def, e := File(ctx.d.Path(), path)
	if e != nil {
		return e
	}

	ctx.idx.AddDefinition(def)
	ctx.state.deps = append(ctx.state.deps, def)
	return nil
}

func jsRewrite(ctx *context, n *ast.CallExpr) error {
	if len(n.Args) == 0 {
		return nil
	}

	var expr string
	var vars []string
	for i, arg := range n.Args {
		// handle expression first
		if i == 0 {
			lit, ok := n.Args[0].(*ast.BasicLit)
			if !ok {
				return fmt.Errorf("process: expected rewrite expression to have basiclit argument, but got %T", n.Args[0])
			}
			exp, e := strconv.Unquote(lit.Value)
			if e != nil {
				return e
			}
			expr = exp
			continue
		}

		// tack on args when i > 0
		a, e := util.ExprToString(arg)
		if e != nil {
			return e
		}
		vars = append(vars, a)
	}

	ctx.state.rewrite = &rewrite{
		expr: expr,
		vars: vars,
	}

	// omit func decl in any rewritten expression
	ctx.state.omit = true
	return nil
}

func isAsync(def def.Definition) (bool, error) {
	fn, ok := def.(Functioner)
	if !ok {
		return false, nil
	}

	async, e := fn.IsAsync()
	if e != nil {
		return false, e
	}

	return async, nil
}

func unique(defs []def.Definition) []def.Definition {
	u := make([]def.Definition, 0, len(defs))
	m := make(map[string]bool)

	for _, def := range defs {
		if _, ok := m[def.ID()]; !ok {
			m[def.ID()] = true
			u = append(u, def)
		}
	}

	return u
}
