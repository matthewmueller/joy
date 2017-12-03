package htmlallcollection

import "github.com/matthewmueller/joy/js"

// HTMLAllCollection struct
// js:"HTMLAllCollection,omit"
type HTMLAllCollection struct {
}

// Item fn
// js:"item"
func (*HTMLAllCollection) Item(nameOrIndex *string) (i interface{}) {
	js.Rewrite("$_.item($1)", nameOrIndex)
	return i
}

// NamedItem fn
// js:"namedItem"
func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$_.namedItem($1)", name)
	return i
}

// Length prop
// js:"length"
func (*HTMLAllCollection) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
