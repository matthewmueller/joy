package defs

import (
	"strconv"
	"strings"
)

type rewrite struct {
	expr     string
	vars     []int
	variadic bool
}

// Rewrite js.Rewrite(expr, arguments...)
func (r *rewrite) Rewrite(caller string, args []string) (string, error) {
	// map out the replacements
	expr := r.expr

	// handle rewrites inside variadic functions
	if r.variadic {
		last := len(r.vars)
		rest := []string{}
		for i, arg := range args {
			if i >= last-1 {
				rest = append(rest, arg)
				continue
			}

			v := r.vars[i]
			expr = strings.Replace(expr, "$"+strconv.Itoa(v), arg, -1)
		}

		lastv := "[" + strings.Join(rest, ",") + "]"
		expr = strings.Replace(expr, "$"+strconv.Itoa(last), lastv, -1)
		return expr, nil
	}

	// handle any other rewrite
	for i, arg := range args {
		v := r.vars[i]
		expr = strings.Replace(expr, "$"+strconv.Itoa(v), arg, -1)
	}

	// handle the caller
	expr = strings.Replace(expr, "$<", caller, -1)

	return expr, nil
}
