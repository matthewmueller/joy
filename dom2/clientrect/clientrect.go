package clientrect

import "github.com/matthewmueller/golly/js"

// js:"ClientRect,omit"
type ClientRect struct {
}

// Bottom
func (*ClientRect) Bottom() (bottom int) {
	js.Rewrite("$<.Bottom")
	return bottom
}

// Bottom
func (*ClientRect) SetBottom(bottom int) {
	js.Rewrite("$<.Bottom = $1", bottom)
}

// Height
func (*ClientRect) Height() (height float32) {
	js.Rewrite("$<.Height")
	return height
}

// Left
func (*ClientRect) Left() (left int) {
	js.Rewrite("$<.Left")
	return left
}

// Left
func (*ClientRect) SetLeft(left int) {
	js.Rewrite("$<.Left = $1", left)
}

// Right
func (*ClientRect) Right() (right int) {
	js.Rewrite("$<.Right")
	return right
}

// Right
func (*ClientRect) SetRight(right int) {
	js.Rewrite("$<.Right = $1", right)
}

// Top
func (*ClientRect) Top() (top int) {
	js.Rewrite("$<.Top")
	return top
}

// Top
func (*ClientRect) SetTop(top int) {
	js.Rewrite("$<.Top = $1", top)
}

// Width
func (*ClientRect) Width() (width float32) {
	js.Rewrite("$<.Width")
	return width
}
