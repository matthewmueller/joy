package fmt

import "github.com/matthewmueller/joy/macro"

// Println formats using the default formats for its operands and writes to standard output.
// Spaces are always added between operands and a newline is appended.
// It returns the number of bytes written and any write error encountered.
func Println(a ...interface{}) (n int, err error) {
	macro.Rewrite("console.log.apply(console.log, $1)", a...)
	return 0, nil
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
//
// NOTE: This differs slightly from the Go implementation because there's no console
// API for printing to the console on the same line. Should be fine though.
func Printf(format string, a ...interface{}) (n int, err error) {
	macro.Rewrite("console.log.apply(console.log, [$1].concat($2))", format, a)
	return 0, nil
}
