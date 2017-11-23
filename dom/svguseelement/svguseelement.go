package svguseelement

import (
	"github.com/matthewmueller/golly/dom/svganimatedlength"
	"github.com/matthewmueller/golly/dom/svganimatedstring"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// SVGUseElement struct
// js:"SVGUseElement,omit"
type SVGUseElement struct {
	window.SVGGraphicsElement
}

// Href prop
func (*SVGUseElement) Href() (href *svganimatedstring.SVGAnimatedString) {
	js.Rewrite("$<.href")
	return href
}

// AnimatedInstanceRoot prop
func (*SVGUseElement) AnimatedInstanceRoot() (animatedInstanceRoot *SVGElementInstance) {
	js.Rewrite("$<.animatedInstanceRoot")
	return animatedInstanceRoot
}

// Height prop
func (*SVGUseElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.height")
	return height
}

// InstanceRoot prop
func (*SVGUseElement) InstanceRoot() (instanceRoot *SVGElementInstance) {
	js.Rewrite("$<.instanceRoot")
	return instanceRoot
}

// Width prop
func (*SVGUseElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.width")
	return width
}

// X prop
func (*SVGUseElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.x")
	return x
}

// Y prop
func (*SVGUseElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.y")
	return y
}
