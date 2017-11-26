package window

import "github.com/matthewmueller/golly/js"

// js:"CSSRule,omit"
type CSSRule interface {

	// CSSText prop
	// js:"cssText,rewrite=csstext"
	CSSText() (cssText string)

	// CSSText prop
	// js:"setcssText,rewrite=setcsstext"
	SetCSSText(cssText string)

	// ParentRule prop
	// js:"parentRule,rewrite=parentrule"
	ParentRule() (parentRule CSSRule)

	// ParentStyleSheet prop
	// js:"parentStyleSheet,rewrite=parentstylesheet"
	ParentStyleSheet() (parentStyleSheet *CSSStyleSheet)

	// Type prop
	// js:"type,rewrite=kind"
	Type() (kind uint8)
}

// csstext prop
func csstext() (cssText string) {
	js.Rewrite("$<.cssText")
	return cssText
}

// setcsstext prop
func setcsstext(cssText string) {
	js.Rewrite("$<.cssText = cssText")
}

// parentrule prop
func parentrule() (parentRule CSSRule) {
	js.Rewrite("$<.parentRule")
	return parentRule
}

// parentstylesheet prop
func parentstylesheet() (parentStyleSheet *CSSStyleSheet) {
	js.Rewrite("$<.parentStyleSheet")
	return parentStyleSheet
}

// kind prop
func kind() (kind uint8) {
	js.Rewrite("$<.type")
	return kind
}
