package domtokenlist

// DOMTokenList interface
// js:"DOMTokenList"
type DOMTokenList interface {

	// Add
	// js:"add"
	// jsrewrite:"$_.add($1)"
	Add(token string)

	// Contains
	// js:"contains"
	// jsrewrite:"$_.contains($1)"
	Contains(token string) (b bool)

	// Item
	// js:"item"
	// jsrewrite:"$_.item($1)"
	Item(index uint) (s string)

	// Remove
	// js:"remove"
	// jsrewrite:"$_.remove($1)"
	Remove(token string)

	// Toggle
	// js:"toggle"
	// jsrewrite:"$_.toggle($1, $2)"
	Toggle(token string, force *bool) (b bool)

	// ToString
	// js:"toString"
	// jsrewrite:"$_.toString()"
	ToString() (s string)

	// Length prop
	// js:"length"
	// jsrewrite:"$_.length"
	Length() (length uint)
}
