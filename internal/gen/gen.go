package gen

import (
	"bytes"
	"text/template"
)

// File that's been generated
type File struct {
	Name   string
	Source string
}

// Vartype struct
type Vartype struct {
	Var      string
	Type     string
	Optional bool
}

// Generate function
func Generate(name string, data interface{}, tpl string) (string, error) {
	t, err := template.New(name).Funcs(funcs()).Parse(tpl)
	if err != nil {
		return "", err
	}

	var b bytes.Buffer
	if e := t.Execute(&b, data); e != nil {
		return "", e
	}

	return string(b.Bytes()), nil
}

// IsBuiltin checks if the name is builtin
func IsBuiltin(name string) bool {
	if builtins[name] == "" {
		return false
	}
	return true
}
