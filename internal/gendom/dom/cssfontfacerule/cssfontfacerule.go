package cssfontfacerule

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/cssrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/cssstyledeclaration"
	"github.com/matthewmueller/golly/js"
)

type CSSFontFaceRule struct {
	*cssrule.CSSRule
}

func (*CSSFontFaceRule) GetStyle() (style *cssstyledeclaration.CSSStyleDeclaration) {
	js.Rewrite("$<.style")
	return style
}
