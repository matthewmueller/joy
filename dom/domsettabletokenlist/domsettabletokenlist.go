package domsettabletokenlist

import (
	"github.com/matthewmueller/golly/dom2/domtokenlist"
	"github.com/matthewmueller/golly/js"
)

// DOMSettableTokenList struct
// js:"DOMSettableTokenList,omit"
type DOMSettableTokenList struct {
	domtokenlist.DOMTokenList
}

// Value prop
func (*DOMSettableTokenList) Value() (value string) {
	js.Rewrite("$<.value")
	return value
}

// Value prop
func (*DOMSettableTokenList) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
