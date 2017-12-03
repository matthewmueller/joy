package strconv

import "github.com/matthewmueller/joy/macro"

// Itoa is shorthand for FormatInt(int64(i), 10).
func Itoa(i int) string {
	macro.Rewrite("String($1)", i)
	return ""
}
