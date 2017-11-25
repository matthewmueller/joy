package defs

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"

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
	rewrite def.Rewrite
	fields  []*field
	params  []string
	rename  string
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
		case *ast.StructType:
			err = structType(ctx, t)
		case *ast.InterfaceType:
			err = interfaceType(ctx, t)
		case *ast.CallExpr:
			err = callExpr(ctx, t)
		case *ast.SelectorExpr:
			err = selectorExpr(ctx, t)
		case *ast.ChanType:
			err = chanType(ctx, t)
		case *ast.Ident:
			err = ident(ctx, t)
		case *ast.CompositeLit:
			err = compositLit(ctx, t)
		}
		return err == nil
	})
	if err != nil {
		return nil, err
	}

	return st, nil
}

func funcDecl(ctx *context, n *ast.FuncDecl) error {
	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return e
	}

	// only update state for these top-level definitions
	switch ctx.d.Kind() {
	case "FUNCTION", "METHOD":
		ctx.state.tag = tag
	}

	if tag != nil && tag.HasOption("async") {
		ctx.state.async = true
	}

	if e := maybeVDOMFuncDecl(ctx, n); e != nil {
		return e
	}

	return nil
}

func genDecl(ctx *context, n *ast.GenDecl) error {
	tag, e := util.JSTag(n.Doc)
	if e != nil {
		return e
	}

	// only update state for these top-level definitions
	switch ctx.d.Kind() {
	case "STRUCT", "VALUE":
		ctx.state.tag = tag
	}

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
			fld.embedded = true
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

	// stct, ok := ctx.d.(Structer)
	// if !ok {
	// 	return errors.New("process/structType: expected struct type to be a structer")
	// }

	// ifaces, err := stct.Implements()
	// if err != nil {
	// 	return err
	// }
	// for _, iface := range ifaces {
	// 	log.Infof("got iface: %s", iface.ID())
	// }

	// for _, method := range stct.Methods() {
	// 	log.Infof("method=%s", method)
	// }

	return nil
}

func interfaceType(ctx *context, n *ast.InterfaceType) error {
	if iface, ok := ctx.d.(*interfaces); ok {
		for _, method := range iface.methods {
			if method.tag == nil {
				continue
			}

			// support js:"name,rewrite=fn"
			for _, option := range method.tag.Options {
				kv := strings.SplitN(option, "=", 2)
				if len(kv) != 2 || kv[0] != "rewrite" {
					continue
				}

				// find the definition
				def := ctx.idx.Get(ctx.d.Path() + " " + kv[1])
				if def == nil {
					return fmt.Errorf("no rewrite function found for %s rewrite=%s", iface.ID(), kv[1])
				}

				// found a rewrite function we'll use
				method.rewriteFunction = def

				// add dependencies
				ctx.state.deps = append(ctx.state.deps, def)
			}
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
	case "js.Runtime":
		return jsRuntime(ctx, n)
	case "js.RawFile":
		return jsFile(ctx, n)
	case "js.Rewrite":
		return jsRewrite(ctx, n)
	case "js.Raw":
		return jsRaw(ctx, n)
	case "vdom.Use":
		return vdomUse(ctx, n)
	}

	def, e := ctx.idx.DefinitionOf(ctx.d.Path(), n)
	if e != nil {
		return e
	}

	switch t := def.(type) {
	case Interfacer:
		m, ok := n.Fun.(*ast.SelectorExpr)
		if !ok {
			return fmt.Errorf("process/callExpr: expected *ast.SelectorExpr but got %T", n.Fun)
		}

		method := t.FindMethod(m.Sel.Name)
		methods := t.ImplementedBy(m.Sel.Name)
		for _, def := range methods {
			ctx.state.deps = append(ctx.state.deps, def)
			if method.Name() != def.Name() {
				return fmt.Errorf("interface must be named the same as all named methods. %s.%s != %s.%s", t.Name(), method.Name(), def.Recv().Name(), def.Name())
			}
		}
	}

	return nil
}

func chanType(ctx *context, n *ast.ChanType) error {
	deps, e := ctx.idx.Runtime("Deferred", "Channel", "Send", "Recv")
	if e != nil {
		return e
	}

	if len(deps) > 0 {
		for _, def := range deps {
			log.Debugf("%s -> %s", ctx.d.ID(), def.ID())
			ctx.state.deps = append(ctx.state.deps, def)
		}
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
		var ignore bool

		// methods shouldn't automatically add their receivers
		method, ok := ctx.d.(Methoder)
		ignore = ok && method.Recv().ID() == def.ID()

		if !ignore {
			log.Debugf("%s -> %s", ctx.d.ID(), def.ID())
			ctx.state.deps = append(ctx.state.deps, def)
		} else {
			log.Debugf("ignoring %s", def.ID())
		}
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
		// log.Infof("def=%s", def.ID())
		var ignore bool

		// methods shouldn't automatically add their receivers
		method, ok := ctx.d.(Methoder)
		ignore = ok && method.Recv().ID() == def.ID()

		if !ignore {
			log.Debugf("%s -> %s", ctx.d.ID(), def.ID())
			ctx.state.deps = append(ctx.state.deps, def)
		}
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

func compositLit(ctx *context, n *ast.CompositeLit) error {
	if e := maybeVDOMCompositLit(ctx, n); e != nil {
		return e
	}

	return nil
}

func jsRuntime(ctx *context, n *ast.CallExpr) error {
	if len(n.Args) == 0 {
		return nil
	}

	var deps []string
	for _, arg := range n.Args {
		lit, ok := arg.(*ast.BasicLit)
		if !ok {
			return fmt.Errorf("fn process: expected rawfile to have basiclit argument, but got %T", arg)
		}

		dep, e := strconv.Unquote(lit.Value)
		if e != nil {
			return e
		}

		deps = append(deps, dep)
	}

	defs, e := ctx.idx.Runtime(deps...)
	if e != nil {
		return e
	}
	if len(defs) == 0 {
		return nil
	}

	// add the runtime path
	runtimePath, e := util.RuntimePath()
	if e != nil {
		return e
	}
	ctx.state.imports["runtime"] = runtimePath

	// update the deps
	for _, def := range defs {
		log.Debugf("%s -> %s", ctx.d.ID(), def.ID())
		ctx.state.deps = append(ctx.state.deps, def)
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

	def, e := File(ctx.d.Path(), path, false)
	if e != nil {
		return e
	}

	ctx.idx.AddDefinition(def)
	log.Debugf("%s -> %s", ctx.d.ID(), def.ID())
	ctx.state.deps = append(ctx.state.deps, def)
	return nil
}

func jsRewrite(ctx *context, n *ast.CallExpr) error {
	if len(n.Args) == 0 {
		return nil
	}

	// if it's a function/method, we'll also want to
	// check if it has variadic arguments
	var variadic bool
	fn, isFn := ctx.d.(Functioner)
	if isFn {
		variadic = fn.IsVariadic()
	}

	var expr string
	var vars []def.RewriteVariable
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

		def, e := ctx.idx.DefinitionOf(ctx.d.Path(), arg)
		if e != nil {
			return e
		} else if def == nil {
			def = ctx.d
		}

		vars = append(vars, &variable{
			def:  def,
			node: arg,
		})
	}

	// async/await support within the rewrite
	if strings.Contains(expr, "async") ||
		strings.Contains(expr, "await") {
		ctx.state.async = true
	}

	ctx.state.rewrite = &rewrite{
		def:      ctx.d,
		expr:     expr,
		vars:     vars,
		variadic: variadic,
	}

	// omit func decl in any rewritten expression
	ctx.state.omit = true
	return nil
}

// right now, just used to figure out if function should be async
func jsRaw(ctx *context, n *ast.CallExpr) error {
	if len(n.Args) == 0 {
		return nil
	}

	var expr string
	for i := range n.Args {
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
	}

	// async/await support within the rewrite
	if strings.Contains(expr, "async") ||
		strings.Contains(expr, "await") {
		ctx.state.async = true
	}

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
