package svgrect

import "github.com/matthewmueller/joy/macro"

// SVGRect struct
// js:"SVGRect,omit"
type SVGRect struct {
}

// Height prop
// js:"height"
func (*SVGRect) Height() (height float32) {
	macro.Rewrite("$_.height")
	return height
}

// SetHeight prop
// js:"height"
func (*SVGRect) SetHeight(height float32) {
	macro.Rewrite("$_.height = $1", height)
}

// Width prop
// js:"width"
func (*SVGRect) Width() (width float32) {
	macro.Rewrite("$_.width")
	return width
}

// SetWidth prop
// js:"width"
func (*SVGRect) SetWidth(width float32) {
	macro.Rewrite("$_.width = $1", width)
}

// X prop
// js:"x"
func (*SVGRect) X() (x float32) {
	macro.Rewrite("$_.x")
	return x
}

// SetX prop
// js:"x"
func (*SVGRect) SetX(x float32) {
	macro.Rewrite("$_.x = $1", x)
}

// Y prop
// js:"y"
func (*SVGRect) Y() (y float32) {
	macro.Rewrite("$_.y")
	return y
}

// SetY prop
// js:"y"
func (*SVGRect) SetY(y float32) {
	macro.Rewrite("$_.y = $1", y)
}
