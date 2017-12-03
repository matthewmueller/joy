package domstringlist

import "github.com/matthewmueller/joy/js"

// DOMStringList struct
// js:"DOMStringList,omit"
type DOMStringList struct {
}

// Contains fn
// js:"contains"
func (*DOMStringList) Contains(str string) (b bool) {
	js.Rewrite("$_.contains($1)", str)
	return b
}

// Item fn
// js:"item"
func (*DOMStringList) Item(index uint) (s string) {
	js.Rewrite("$_.item($1)", index)
	return s
}

// Length prop
// js:"length"
func (*DOMStringList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
