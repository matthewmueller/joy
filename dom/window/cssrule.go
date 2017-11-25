package window

// js:"CSSRule,omit"
type CSSRule interface {

	// CSSText prop
	// js:"cssText"
	CSSText() (cssText string)

	// CSSText prop
	// js:"setcssText"
	SetCSSText(cssText string)

	// ParentRule prop
	// js:"parentRule"
	ParentRule() (parentRule CSSRule)

	// ParentStyleSheet prop
	// js:"parentStyleSheet"
	ParentStyleSheet() (parentStyleSheet *CSSStyleSheet)

	// Type prop
	// js:"type"
	Type() (kind uint8)
}
