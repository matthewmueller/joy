package defs

import (
	"strconv"
	"strings"
)

type rewrite struct {
	kind string
	expr string
	vars []int
}

// Rewrite js.Rewrite(expr, arguments...)
func (r *rewrite) Rewrite(args []string) (string, error) {
	// map out the replacements
	// replacements := map[string]string{}
	expr := r.expr
	for i, arg := range args {
		v := r.vars[i]
		expr = strings.Replace(expr, "$"+strconv.Itoa(v), arg, -1)
	}

	return expr, nil
}
