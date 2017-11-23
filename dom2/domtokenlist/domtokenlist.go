package domtokenlist

// js:"DOMTokenList,omit"
type DOMTokenList interface {

	// Add
	Add(token string)

	// Contains
	Contains(token string) (b bool)

	// Item
	Item(index uint) (s string)

	// Remove
	Remove(token string)

	// Toggle
	Toggle(token string, force *bool) (b bool)

	// ToString
	ToString() (s string)

	// Length
	Length() (length uint)
}
