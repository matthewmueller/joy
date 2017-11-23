package svgrect

import "github.com/matthewmueller/golly/js"

// SVGRect struct
// js:"SVGRect,omit"
type SVGRect struct {
}

// Height
func (*SVGRect) Height() (height float32) {
	js.Rewrite("$<.Height")
	return height
}

// Height
func (*SVGRect) SetHeight(height float32) {
	js.Rewrite("$<.Height = $1", height)
}

// Width
func (*SVGRect) Width() (width float32) {
	js.Rewrite("$<.Width")
	return width
}

// Width
func (*SVGRect) SetWidth(width float32) {
	js.Rewrite("$<.Width = $1", width)
}

// X
func (*SVGRect) X() (x float32) {
	js.Rewrite("$<.X")
	return x
}

// X
func (*SVGRect) SetX(x float32) {
	js.Rewrite("$<.X = $1", x)
}

// Y
func (*SVGRect) Y() (y float32) {
	js.Rewrite("$<.Y")
	return y
}

// Y
func (*SVGRect) SetY(y float32) {
	js.Rewrite("$<.Y = $1", y)
}
