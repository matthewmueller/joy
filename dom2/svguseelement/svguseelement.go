package svguseelement

import "github.com/matthewmueller/golly/js"

// js:"SVGUseElement,omit"
type SVGUseElement struct {
	window.SVGGraphicsElement
	svgurireference.SVGURIReference
}

// AnimatedInstanceRoot
func (*SVGUseElement) AnimatedInstanceRoot() (animatedInstanceRoot *SVGElementInstance) {
	js.Rewrite("$<.AnimatedInstanceRoot")
	return animatedInstanceRoot
}

// Height
func (*SVGUseElement) Height() (height *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Height")
	return height
}

// InstanceRoot
func (*SVGUseElement) InstanceRoot() (instanceRoot *SVGElementInstance) {
	js.Rewrite("$<.InstanceRoot")
	return instanceRoot
}

// Width
func (*SVGUseElement) Width() (width *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Width")
	return width
}

// X
func (*SVGUseElement) X() (x *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.X")
	return x
}

// Y
func (*SVGUseElement) Y() (y *svganimatedlength.SVGAnimatedLength) {
	js.Rewrite("$<.Y")
	return y
}
