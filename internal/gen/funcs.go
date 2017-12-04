package gen

import (
	"strconv"
	"strings"
	"text/template"

	"github.com/knq/snaker"
)

func funcs() template.FuncMap {
	return template.FuncMap{
		"capitalize": Capitalize,
		"lowercase":  Lowercase,
		"pointer":    Pointer,
		"variable":   Variable,
		"joinvt":     JoinVT,
		"joinv":      JoinV,
		"join":       Join,
		"vt":         VT,
		"sequence":   Sequence,
		"identifier": Identifier,
	}
}

// Identifier from the string avoiding the builtins
func Identifier(s string) string {
	out := []string{}
	for i, w := range strings.Split(s, "-") {
		if i == 0 {
			out = append(out, w)
			continue
		}

		capitalize := strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		out = append(out, capitalize)
	}
	s = strings.Join(out, "")

	out = []string{}
	for i, w := range strings.Split(s, "_") {
		if i == 0 {
			out = append(out, w)
			continue
		}

		capitalize := strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		out = append(out, capitalize)
	}
	s = strings.Join(out, "")

	if builtins[s] != "" {
		return builtins[s]
	}
	return s
}

// Capitalize helper
func Capitalize(s string) string {
	s = snaker.SnakeToCamelIdentifier(snaker.CamelToSnake(s))
	return Identifier(s)
}

// Lowercase fn
func Lowercase(s string) string {
	s = strings.ToLower(snaker.SnakeToCamelIdentifier(s))
	return Identifier(s)
}

// Variable creates a short variable from the string
func Variable(s string) string {
	s = strings.TrimLeft(s, "*[]")
	if len(s) == 0 {
		return s
	}
	letter := s[:1]
	return Lowercase(letter)
}

// Pointer returns a pointer version of the variable
func Pointer(s string) string {
	if zero, isset := defaults[s]; isset && zero == nil {
		return s
	}
	return "*" + strings.TrimLeft(s, "*")
}

// VT turns a vartype into a string
func VT(vt Vartype) string {
	name := Identifier(vt.Var)
	if vt.Optional {
		return name + " " + Pointer(vt.Type)
	}
	return name + " " + vt.Type
}

// JoinVT joins vartypes var and type
func JoinVT(vts []Vartype) string {
	var a []string
	for _, vt := range vts {
		a = append(a, VT(vt))
	}
	return strings.Join(a, ", ")
}

// JoinV joins vartypes just by the var
func JoinV(vts []Vartype) string {
	var a []string
	for _, vt := range vts {
		a = append(a, Identifier(vt.Var))
	}
	return strings.Join(a, ", ")
}

// Join just joins 2 strings by commas
func Join(a []string) string {
	return strings.Join(a, ", ")
}

// Sequence counts up until a number
// starting at $1
func Sequence(until int) []string {
	var n []string
	for i := 0; i < until; i++ {
		n = append(n, "$"+strconv.Itoa(i+1))
	}
	return n
}
