package window

// js:"CSSRule,omit"
type CSSRule interface {

	// CSSText
	CSSText() (cssText string)

	// CSSText
	SetCSSText(cssText string)

	// ParentRule
	ParentRule() (parentRule CSSRule)

	// ParentStyleSheet
	ParentStyleSheet() (parentStyleSheet *CSSStyleSheet)

	// Type
	Type() (kind uint8)
}
