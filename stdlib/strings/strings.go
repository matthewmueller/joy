package strings

import "github.com/matthewmueller/golly/js"

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
	js.Rewrite("$1.join($2)", a, sep)
	return ""
}
