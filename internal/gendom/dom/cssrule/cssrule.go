package cssrule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstylesheet"
	"github.com/matthewmueller/golly/js"
)

type CSSRule struct {
}

func (*CSSRule) GetCSSText() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

func (*CSSRule) SetCSSText(cssText string) {
	js.Rewrite("$<.cssText = $1", cssText)
}

func (*CSSRule) GetParentRule() (parentRule *CSSRule) {
	js.Rewrite("$<.parentRule")
	return parentRule
}

func (*CSSRule) GetParentStyleSheet() (parentStyleSheet *cssstylesheet.CSSStyleSheet) {
	js.Rewrite("$<.parentStyleSheet")
	return parentStyleSheet
}

func (*CSSRule) GetType() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}
