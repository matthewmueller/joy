package domsettabletokenlist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/domtokenlist"
	"github.com/matthewmueller/golly/js"
)

type DOMSettableTokenList struct {
	*domtokenlist.DOMTokenList
}

func (*DOMSettableTokenList) GetValue() (value string) {
	js.Rewrite("$<.value")
	return value
}

func (*DOMSettableTokenList) SetValue(value string) {
	js.Rewrite("$<.value = $1", value)
}
