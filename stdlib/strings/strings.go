package strings

import (
	"github.com/matthewmueller/joy/macro"
)

// Join concatenates the elements of a to create a single string. The separator string
// sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
	macro.Rewrite("$1.join($2)", a, sep)
	return ""
}

// ToUpper returns a copy of the string s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	macro.Rewrite("$1.toUpperCase()", s)
	return ""
}

// ToLower returns a copy of the string s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	macro.Rewrite("$1.toLowerCase()", s)
	return ""
}
