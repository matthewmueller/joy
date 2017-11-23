package domstringlist

import "github.com/matthewmueller/golly/js"

// DOMStringList struct
// js:"DOMStringList,omit"
type DOMStringList struct {
}

// Contains fn
func (*DOMStringList) Contains(str string) (b bool) {
	js.Rewrite("$<.contains($1)", str)
	return b
}

// Item fn
func (*DOMStringList) Item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

// Length prop
func (*DOMStringList) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
