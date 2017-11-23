package htmlallcollection

import "github.com/matthewmueller/golly/js"

// HTMLAllCollection struct
// js:"HTMLAllCollection,omit"
type HTMLAllCollection struct {
}

// Item
func (*HTMLAllCollection) Item(nameOrIndex *string) (i interface{}) {
	js.Rewrite("$<.Item($1)", nameOrIndex)
	return i
}

// NamedItem
func (*HTMLAllCollection) NamedItem(name string) (i interface{}) {
	js.Rewrite("$<.NamedItem($1)", name)
	return i
}

// Length
func (*HTMLAllCollection) Length() (length uint) {
	js.Rewrite("$<.Length")
	return length
}
