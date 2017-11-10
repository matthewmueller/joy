package strconv

import "github.com/matthewmueller/golly/js"

// Itoa is shorthand for FormatInt(int64(i), 10).
func Itoa(i int) string {
	js.Rewrite("String($1)", i)
	return ""
}
