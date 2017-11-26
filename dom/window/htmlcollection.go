package window

import "github.com/matthewmueller/golly/js"

// js:"HTMLCollection,omit"
type HTMLCollection interface {

	// Item Retrieves an object from various collections.
	// js:"item,rewrite=item"
	Item(index uint) (e Element)

	// NamedItem Retrieves a select object or an object from an options collection.
	// js:"namedItem,rewrite=nameditem"
	NamedItem(name string) (e Element)

	// Length prop Sets or retrieves the number of objects in a collection.
	// js:"length,rewrite=length"
	Length() (length uint)
}

// item fn Retrieves an object from various collections.
func item(index uint) (e Element) {
	js.Rewrite("$<.item($1)", index)
	return e
}

// nameditem fn Retrieves a select object or an object from an options collection.
func nameditem(name string) (e Element) {
	js.Rewrite("$<.namedItem($1)", name)
	return e
}

// // length prop Sets or retrieves the number of objects in a collection.
// func length() (length uint) {
// 	js.Rewrite("$<.length")
// 	return length
// }
