package domsettabletokenlist

import (
	"github.com/matthewmueller/joy/dom/domtokenlist"
	"github.com/matthewmueller/joy/js"
)

var _ domtokenlist.DOMTokenList = (*DOMSettableTokenList)(nil)

// DOMSettableTokenList struct
// js:"DOMSettableTokenList,omit"
type DOMSettableTokenList struct {
}

// Add fn
// js:"add"
func (*DOMSettableTokenList) Add(token string) {
	js.Rewrite("$_.add($1)", token)
}

// Contains fn
// js:"contains"
func (*DOMSettableTokenList) Contains(token string) (b bool) {
	js.Rewrite("$_.contains($1)", token)
	return b
}

// Item fn
// js:"item"
func (*DOMSettableTokenList) Item(index uint) (s string) {
	js.Rewrite("$_.item($1)", index)
	return s
}

// Remove fn
// js:"remove"
func (*DOMSettableTokenList) Remove(token string) {
	js.Rewrite("$_.remove($1)", token)
}

// Toggle fn
// js:"toggle"
func (*DOMSettableTokenList) Toggle(token string, force *bool) (b bool) {
	js.Rewrite("$_.toggle($1, $2)", token, force)
	return b
}

// ToString fn
// js:"toString"
func (*DOMSettableTokenList) ToString() (s string) {
	js.Rewrite("$_.toString()")
	return s
}

// Value prop
// js:"value"
func (*DOMSettableTokenList) Value() (value string) {
	js.Rewrite("$_.value")
	return value
}

// SetValue prop
// js:"value"
func (*DOMSettableTokenList) SetValue(value string) {
	js.Rewrite("$_.value = $1", value)
}

// Length prop
// js:"length"
func (*DOMSettableTokenList) Length() (length uint) {
	js.Rewrite("$_.length")
	return length
}
