package svgtransformlist

import (
	"github.com/matthewmueller/joy/dom/svgmatrix"
	"github.com/matthewmueller/joy/dom/svgtransform"
	"github.com/matthewmueller/joy/js"
)

// SVGTransformList struct
// js:"SVGTransformList,omit"
type SVGTransformList struct {
}

// AppendItem fn
// js:"appendItem"
func (*SVGTransformList) AppendItem(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.appendItem($1)", newItem)
	return s
}

// Clear fn
// js:"clear"
func (*SVGTransformList) Clear() {
	js.Rewrite("$_.clear()")
}

// Consolidate fn
// js:"consolidate"
func (*SVGTransformList) Consolidate() (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.consolidate()")
	return s
}

// CreateSVGTransformFromMatrix fn
// js:"createSVGTransformFromMatrix"
func (*SVGTransformList) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.createSVGTransformFromMatrix($1)", matrix)
	return s
}

// GetItem fn
// js:"getItem"
func (*SVGTransformList) GetItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.getItem($1)", index)
	return s
}

// Initialize fn
// js:"initialize"
func (*SVGTransformList) Initialize(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.initialize($1)", newItem)
	return s
}

// InsertItemBefore fn
// js:"insertItemBefore"
func (*SVGTransformList) InsertItemBefore(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.insertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem fn
// js:"removeItem"
func (*SVGTransformList) RemoveItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.removeItem($1)", index)
	return s
}

// ReplaceItem fn
// js:"replaceItem"
func (*SVGTransformList) ReplaceItem(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$_.replaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems prop
// js:"numberOfItems"
func (*SVGTransformList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$_.numberOfItems")
	return numberOfItems
}
