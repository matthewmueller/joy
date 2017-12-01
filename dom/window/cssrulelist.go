package window

import "github.com/matthewmueller/golly/js"

// CSSRuleList struct
// js:"CSSRuleList,omit"
type CSSRuleList struct {
}

// Item fn
// js:"item"
func (*CSSRuleList) Item(index uint) (c CSSRule) {
	js.Rewrite("$_.item($1)", index)
	return c
}

// Length prop
// js:"length"
func (*CSSRuleList) Length() (length int) {
	js.Rewrite("$_.length")
	return length
}
