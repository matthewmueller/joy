package imagedata

import "github.com/matthewmueller/joy/macro"

// New fn
func New(width uint, height uint) *ImageData {
	macro.Rewrite("ImageData")
	return &ImageData{}
}

// ImageData struct
// js:"ImageData,omit"
type ImageData struct {
}

// Data prop
// js:"data"
func (*ImageData) Data() (data []uint8) {
	macro.Rewrite("$_.data")
	return data
}

// Height prop
// js:"height"
func (*ImageData) Height() (height uint) {
	macro.Rewrite("$_.height")
	return height
}

// Width prop
// js:"width"
func (*ImageData) Width() (width uint) {
	macro.Rewrite("$_.width")
	return width
}
