package htmlallcollection

import "github.com/matthewmueller/joy/macro"

// HTMLAllCollection struct
// js:"HTMLAllCollection,omit"
type HTMLAllCollection struct {
}

// Item fn
// js:"item"
func (*HTMLAllCollection) Item(nameOrIndex *string) (i interface{}) {
	macro.Rewrite("$_.item($1)", nameOrIndex)
	return i
}

// NamedItem fn
// js:"namedItem"
func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	macro.Rewrite("$_.namedItem($1)", name)
	return i
}

// Length prop
// js:"length"
func (*HTMLAllCollection) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}
