package window

// CSSRule interface
// js:"CSSRule"
type CSSRule interface {

	// CSSText prop
	// js:"cssText"
	// jsrewrite:"$_.cssText"
	CSSText() (cssText string)

	// SetCSSText prop
	// js:"cssText"
	// jsrewrite:"$_.cssText = $1"
	SetCSSText(cssText string)

	// ParentRule prop
	// js:"parentRule"
	// jsrewrite:"$_.parentRule"
	ParentRule() (parentRule CSSRule)

	// ParentStyleSheet prop
	// js:"parentStyleSheet"
	// jsrewrite:"$_.parentStyleSheet"
	ParentStyleSheet() (parentStyleSheet *CSSStyleSheet)

	// Type prop
	// js:"type"
	// jsrewrite:"$_.type"
	Type() (kind uint8)
}
