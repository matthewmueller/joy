package defs

import (
	"go/ast"

	"github.com/matthewmueller/joy/internal/compiler/def"
	"github.com/matthewmueller/joy/internal/compiler/util"
)

var _ def.Rewrite = (*rewrite)(nil)
var _ def.RewriteVariable = (*variable)(nil)

type rewrite struct {
	rewritee def.Rewritee
	expr     string
	vars     []def.RewriteVariable // variables passed into js.Rewrite(expr, <variables>)
}

func (r *rewrite) Rewritee() def.Rewritee {
	return r.rewritee
}

func (r *rewrite) Expression() string {
	return r.expr
}

func (r *rewrite) Vars() []def.RewriteVariable {
	return r.vars
}

type variable struct {
	def  def.Definition
	node ast.Expr
}

func (v *variable) Definition() def.Definition {
	return v.def
}

func (v *variable) Node() ast.Expr {
	return v.node
}

func (v *variable) String() (string, error) {
	return util.ExprToString(v.node)
}
