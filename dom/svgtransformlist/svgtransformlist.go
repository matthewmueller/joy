package svgtransformlist

import (
	"github.com/matthewmueller/golly/dom2/svgmatrix"
	"github.com/matthewmueller/golly/dom2/svgtransform"
	"github.com/matthewmueller/golly/js"
)

// SVGTransformList struct
// js:"SVGTransformList,omit"
type SVGTransformList struct {
}

// AppendItem fn
func (*SVGTransformList) AppendItem(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.appendItem($1)", newItem)
	return s
}

// Clear fn
func (*SVGTransformList) Clear() {
	js.Rewrite("$<.clear()")
}

// Consolidate fn
func (*SVGTransformList) Consolidate() (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.consolidate()")
	return s
}

// CreateSVGTransformFromMatrix fn
func (*SVGTransformList) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.createSVGTransformFromMatrix($1)", matrix)
	return s
}

// GetItem fn
func (*SVGTransformList) GetItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.getItem($1)", index)
	return s
}

// Initialize fn
func (*SVGTransformList) Initialize(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
func (*SVGTransformList) InsertItemBefore(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
func (*SVGTransformList) RemoveItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.removeItem($1)", index)
	return s
}

// ReplaceItem fn
func (*SVGTransformList) ReplaceItem(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
func (*SVGTransformList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.numberOfItems")
	return numberOfItems
}
