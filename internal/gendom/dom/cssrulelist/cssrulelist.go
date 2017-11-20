package cssrulelist

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/js"
)

type CSSRuleList struct {
}

func (*CSSRuleList) Item(index uint) (c *cssrule.CSSRule) {
	js.Rewrite("$<.item($1)", index)
	return c
}

func (*CSSRuleList) GetLength() (length int) {
	js.Rewrite("$<.length")
	return length
}
