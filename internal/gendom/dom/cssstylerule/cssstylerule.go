package cssstylerule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/js"
)

type CSSStyleRule struct {
	*cssrule.CSSRule
}

func (*CSSStyleRule) GetReadOnly() (readOnly bool) {
	js.Rewrite("$<.readOnly")
	return readOnly
}

func (*CSSStyleRule) GetSelectorText() (selectorText string) {
	js.Rewrite("$<.selectorText")
	return selectorText
}

func (*CSSStyleRule) SetSelectorText(selectorText string) {
	js.Rewrite("$<.selectorText = $1", selectorText)
}

func (*CSSStyleRule) GetStyle() (style *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
