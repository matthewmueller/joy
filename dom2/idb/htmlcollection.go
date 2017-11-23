package idb

// js:"HTMLCollection,omit"
type HTMLCollection interface {

	// Item Retrieves an object from various collections.
	Item(index uint) (e Element)

	// NamedItem Retrieves a select object or an object from an options collection.
	NamedItem(name string) (e Element)

	// Length Sets or retrieves the number of objects in a collection.
	Length() (length uint)
}
