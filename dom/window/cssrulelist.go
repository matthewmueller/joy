package window

import "github.com/matthewmueller/golly/js"

// CSSRuleList struct
// js:"CSSRuleList,omit"
type CSSRuleList struct {
}

// Item fn
func (*CSSRuleList) Item(index uint) (c CSSRule) {
	js.Rewrite("$<.item($1)", index)
	return c
}

// Length prop
func (*CSSRuleList) Length() (length int) {
	js.Rewrite("$<.length")
	return length
}
