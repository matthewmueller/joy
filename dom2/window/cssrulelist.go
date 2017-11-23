package window

import "github.com/matthewmueller/golly/js"

// CSSRuleList struct
// js:"CSSRuleList,omit"
type CSSRuleList struct {
}

// Item
func (*CSSRuleList) Item(index uint) (c CSSRule) {
	js.Rewrite("$<.Item($1)", index)
	return c
}

// Length
func (*CSSRuleList) Length() (length int) {
	js.Rewrite("$<.Length")
	return length
}
