package domstringlist

import "github.com/matthewmueller/golly/js"

// DOMStringList struct
// js:"DOMStringList,omit"
type DOMStringList struct {
}

// Contains
func (*DOMStringList) Contains(str string) (b bool) {
	js.Rewrite("$<.Contains($1)", str)
	return b
}

// Item
func (*DOMStringList) Item(index uint) (s string) {
	js.Rewrite("$<.Item($1)", index)
	return s
}

// Length
func (*DOMStringList) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
