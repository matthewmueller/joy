package domtokenlist

// js:"DOMTokenList,omit"
type DOMTokenList interface {

	// Add
	// js:"add"
	Add(token string)

	// Contains
	// js:"contains"
	Contains(token string) (b bool)

	// Item
	// js:"item"
	Item(index uint) (s string)

	// Remove
	// js:"remove"
	Remove(token string)

	// Toggle
	// js:"toggle"
	Toggle(token string, force *bool) (b bool)

	// ToString
	// js:"toString"
	ToString() (s string)

	// Length prop
	// js:"length"
	Length() (length uint)
}
