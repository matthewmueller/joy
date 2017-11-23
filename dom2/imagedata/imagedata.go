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

// Data
func (*ImageData) Data() (data []uint8) {
	js.Rewrite("$<.Data")
	return data
}

// Height
func (*ImageData) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// Width
func (*ImageData) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}
