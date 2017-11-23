package imagedata

import "github.com/matthewmueller/golly/js"

// New fn
func New(width uint, height uint) *ImageData {
	js.Rewrite("ImageData")
	return &ImageData{}
}

// ImageData struct
// js:"ImageData,omit"
type ImageData struct {
}

// Data prop
func (*ImageData) Data() (data []uint8) {
	js.Rewrite("$<.data")
	return data
}

// Height prop
func (*ImageData) Height() (height uint) {
	js.Rewrite("$<.height")
	return height
}

// Width prop
func (*ImageData) Width() (width uint) {
	js.Rewrite("$<.width")
	return width
}
