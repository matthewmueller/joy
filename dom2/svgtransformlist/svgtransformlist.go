package svgtransformlist

import "github.com/matthewmueller/golly/js"

// js:"SVGTransformList,omit"
type SVGTransformList struct {
}

// AppendItem
func (*SVGTransformList) AppendItem(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.AppendItem($1)", newItem)
	return s
}

// Clear
func (*SVGTransformList) Clear() {
	js.Rewrite("$<.Clear()")
}

// Consolidate
func (*SVGTransformList) Consolidate() (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.Consolidate()")
	return s
}

// CreateSVGTransformFromMatrix
func (*SVGTransformList) CreateSVGTransformFromMatrix(matrix *svgmatrix.SVGMatrix) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.CreateSVGTransformFromMatrix($1)", matrix)
	return s
}

// GetItem
func (*SVGTransformList) GetItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.GetItem($1)", index)
	return s
}

// Initialize
func (*SVGTransformList) Initialize(newItem *svgtransform.SVGTransform) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.Initialize($1)", newItem)
	return s
}

// InsertItemBefore
func (*SVGTransformList) InsertItemBefore(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.InsertItemBefore($1, $2)", newItem, index)
	return s
}

// RemoveItem
func (*SVGTransformList) RemoveItem(index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.RemoveItem($1)", index)
	return s
}

// ReplaceItem
func (*SVGTransformList) ReplaceItem(newItem *svgtransform.SVGTransform, index uint) (s *svgtransform.SVGTransform) {
	js.Rewrite("$<.ReplaceItem($1, $2)", newItem, index)
	return s
}

// NumberOfItems
func (*SVGTransformList) NumberOfItems() (numberOfItems uint) {
	js.Rewrite("$<.NumberOfItems")
	return numberOfItems
}
