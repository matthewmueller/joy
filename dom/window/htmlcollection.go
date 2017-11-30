package window

// HTMLCollection interface
// js:"HTMLCollection"
type HTMLCollection interface {

	// Item Retrieves an object from various collections.
	// js:"item"
	Item(index uint) (e Element)

	// NamedItem Retrieves a select object or an object from an options collection.
	// js:"namedItem"
	NamedItem(name string) (e Element)

	// Length prop Sets or retrieves the number of objects in a collection.
	// js:"length"
	Length() (length uint)
}
