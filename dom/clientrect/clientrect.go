package clientrect

import "github.com/matthewmueller/golly/js"

// ClientRect struct
// js:"ClientRect,omit"
type ClientRect struct {
}

// Bottom prop
func (*ClientRect) Bottom() (bottom int) {
	js.Rewrite("$<.bottom")
	return bottom
}

// Bottom prop
func (*ClientRect) SetBottom(bottom int) {
	js.Rewrite("$<.bottom = $1", bottom)
}

// Height prop
func (*ClientRect) Height() (height float32) {
	js.Rewrite("$<.height")
	return height
}

// Left prop
func (*ClientRect) Left() (left int) {
	js.Rewrite("$<.left")
	return left
}

// Left prop
func (*ClientRect) SetLeft(left int) {
	js.Rewrite("$<.left = $1", left)
}

// Right prop
func (*ClientRect) Right() (right int) {
	js.Rewrite("$<.right")
	return right
}

// Right prop
func (*ClientRect) SetRight(right int) {
	js.Rewrite("$<.right = $1", right)
}

// Top prop
func (*ClientRect) Top() (top int) {
	js.Rewrite("$<.top")
	return top
}

// Top prop
func (*ClientRect) SetTop(top int) {
	js.Rewrite("$<.top = $1", top)
}

// Width prop
func (*ClientRect) Width() (width float32) {
	js.Rewrite("$<.width")
	return width
}
