package svgrect

import "github.com/matthewmueller/golly/js"

// SVGRect struct
// js:"SVGRect,omit"
type SVGRect struct {
}

// Height prop
func (*SVGRect) Height() (height float32) {
	js.Rewrite("$<.height")
	return height
}

// Height prop
func (*SVGRect) SetHeight(height float32) {
	js.Rewrite("$<.height = $1", height)
}

// Width prop
func (*SVGRect) Width() (width float32) {
	js.Rewrite("$<.width")
	return width
}

// Width prop
func (*SVGRect) SetWidth(width float32) {
	js.Rewrite("$<.width = $1", width)
}

// X prop
func (*SVGRect) X() (x float32) {
	js.Rewrite("$<.x")
	return x
}

// X prop
func (*SVGRect) SetX(x float32) {
	js.Rewrite("$<.x = $1", x)
}

// Y prop
func (*SVGRect) Y() (y float32) {
	js.Rewrite("$<.y")
	return y
}

// Y prop
func (*SVGRect) SetY(y float32) {
	js.Rewrite("$<.y = $1", y)
}
