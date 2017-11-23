package htmlallcollection

import "github.com/matthewmueller/golly/js"

// HTMLAllCollection struct
// js:"HTMLAllCollection,omit"
type HTMLAllCollection struct {
}

// Item fn
func (*HTMLAllCollection) Item(nameOrIndex *string) (i interface{}) {
	js.Rewrite("$<.item($1)", nameOrIndex)
	return i
}

// NamedItem fn
func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.namedItem($1)", name)
	return i
}

// Length prop
func (*HTMLAllCollection) Length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
