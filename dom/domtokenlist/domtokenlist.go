package domtokenlist

import "github.com/matthewmueller/golly/js"

// js:"DOMTokenList,omit"
type DOMTokenList interface {

	// Add
	// js:"add,rewrite=add"
	Add(token string)

	// Contains
	// js:"contains,rewrite=contains"
	Contains(token string) (b bool)

	// Item
	// js:"item,rewrite=item"
	Item(index uint) (s string)

	// Remove
	// js:"remove,rewrite=remove"
	Remove(token string)

	// Toggle
	// js:"toggle,rewrite=toggle"
	Toggle(token string, force *bool) (b bool)

	// ToString
	// js:"toString,rewrite=tostring"
	ToString() (s string)

	// Length prop
	// js:"length,rewrite=length"
	Length() (length uint)
}

// add fn
func add(token string) {
	js.Rewrite("$<.add($1)", token)
}

// contains fn
func contains(token string) (b bool) {
	js.Rewrite("$<.contains($1)", token)
	return b
}

// item fn
func item(index uint) (s string) {
	js.Rewrite("$<.item($1)", index)
	return s
}

// remove fn
func remove(token string) {
	js.Rewrite("$<.remove($1)", token)
}

// toggle fn
func toggle(token string, force *bool) (b bool) {
	js.Rewrite("$<.toggle($1, $2)", token, force)
	return b
}

// tostring fn
func tostring() (s string) {
	js.Rewrite("$<.toString()")
	return s
}

// length prop
func length() (length uint) {
	js.Rewrite("$<.length")
	return length
}
