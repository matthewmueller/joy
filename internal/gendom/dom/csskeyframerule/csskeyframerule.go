package csskeyframerule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/js"
)

type CSSKeyframeRule struct {
	*cssrule.CSSRule
}

func (*CSSKeyframeRule) GetKeyText() (keyText string) {
	js.Rewrite("$<.keyText")
	return keyText
}

func (*CSSKeyframeRule) SetKeyText(keyText string) {
	js.Rewrite("$<.keyText = $1", keyText)
}

func (*CSSKeyframeRule) GetStyle() (style *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
