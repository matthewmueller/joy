package defs

import (
	"errors"
	"strconv"
	"strings"
)

type rewrite struct {
	expr string
	vars []string
}

// Rewrite js.Rewrite(expr, arguments...)
func (r *rewrite) Rewrite(args []string) (string, error) {
	if len(args) != len(r.vars) {
		return "", errors.New("Rewrite doesn't support spreads (e.g param...) yet")
	}

	// map out the replacements
	replacements := map[string]string{}
	for i, arg := range args {
		replacements[r.vars[i]] = arg
	}

	// make the substitutions
	expr := r.expr
	for i, variable := range r.vars {
		replacement, isset := replacements[variable]
		if !isset {
			return "", errors.New("js.Rewrite() variable doesn't match the function parameter")
		}
		expr = strings.Replace(expr, "$"+strconv.Itoa(i+1), replacement, -1)
	}

	return expr, nil
}
