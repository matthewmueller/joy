package domstringlist

import "github.com/matthewmueller/joy/macro"

// DOMStringList struct
// js:"DOMStringList,omit"
type DOMStringList struct {
}

// Contains fn
// js:"contains"
func (*DOMStringList) Contains(str string) (b bool) {
	macro.Rewrite("$_.contains($1)", str)
	return b
}

// Item fn
// js:"item"
func (*DOMStringList) Item(index uint) (s string) {
	macro.Rewrite("$_.item($1)", index)
	return s
}

// Length prop
// js:"length"
func (*DOMStringList) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
