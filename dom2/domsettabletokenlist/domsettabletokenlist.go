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

// Value
func (*DOMSettableTokenList) Value() (value string) {
	js.Rewrite("$<.Value")
	return value
}

// Value
func (*DOMSettableTokenList) SetValue(value string) {
	js.Rewrite("$<.Value = $1", value)
}
