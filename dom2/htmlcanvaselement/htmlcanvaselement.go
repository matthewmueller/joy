package htmlcanvaselement

import (
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// js:"HTMLCanvasElement,omit"
type HTMLCanvasElement struct {
	window.HTMLElement
}

// GetContext Returns an object that provides methods and properties for drawing and manipulating images and graphics on a canvas element in a document. A context object includes information about colors, line widths, fonts, and other graphic parameters that can be drawn on a canvas.
//     * @param contextId The identifier (ID) of the type of canvas to create. Internet Explorer 9 and Internet Explorer 10 support only a 2-D context using canvas.getContext("2d"); IE11 Preview also supports 3-D or WebGL context using canvas.getContext("experimental-webgl");
func (*HTMLCanvasElement) GetContext(contextId string, args interface{}) (i interface{}) {
	js.Rewrite("$<.GetContext($1, $2)", contextId, args)
	return i
}

// MsToBlob Returns a blob object encoded as a Portable Network Graphics (PNG) format from a canvas image or drawing.
func (*HTMLCanvasElement) MsToBlob() (b blob.Blob) {
	js.Rewrite("$<.MsToBlob()")
	return b
}

// ToDataURL Returns the content of the current canvas as an image that you can use as a source for another canvas or an HTML element.
//     * @param type The standard MIME type for the image format to return. If you do not specify this parameter, the default value is a PNG format image.
func (*HTMLCanvasElement) ToDataURL(kind *string, args interface{}) (s string) {
	js.Rewrite("$<.ToDataURL($1, $2)", kind, args)
	return s
}

// Height Gets or sets the height of a canvas element on a document.
func (*HTMLCanvasElement) Height() (height uint) {
	js.Rewrite("$<.Height")
	return height
}

// Height Gets or sets the height of a canvas element on a document.
func (*HTMLCanvasElement) SetHeight(height uint) {
	js.Rewrite("$<.Height = $1", height)
}

// Width Gets or sets the width of a canvas element on a document.
func (*HTMLCanvasElement) Width() (width uint) {
	js.Rewrite("$<.Width")
	return width
}

// Width Gets or sets the width of a canvas element on a document.
func (*HTMLCanvasElement) SetWidth(width uint) {
	js.Rewrite("$<.Width = $1", width)
}
