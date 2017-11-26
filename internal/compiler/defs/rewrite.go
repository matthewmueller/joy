package defs

import (
	"go/ast"

	"github.com/matthewmueller/golly/internal/compiler/def"
	"github.com/matthewmueller/golly/internal/compiler/util"
)

var _ def.Rewrite = (*rewrite)(nil)
var _ def.RewriteVariable = (*variable)(nil)

type rewrite struct {
	def      def.Definition
	expr     string
	vars     []def.RewriteVariable
	variadic bool
}

func (r *rewrite) Definition() def.Definition {
	return r.def
}

func (r *rewrite) Expression() string {
	return r.expr
}

func (r *rewrite) Vars() []def.RewriteVariable {
	return r.vars
}

func (r *rewrite) Variadic() bool {
	return r.variadic
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
