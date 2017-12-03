package path2d

import "github.com/matthewmueller/joy/js"

// New fn
func New(path *Path2d) *Path2d {
	js.Rewrite("Path2D")
	return &Path2d{}
}

// Path2d struct
// js:"Path2d,omit"
type Path2d struct {
}

// Arc fn
// js:"arc"
func (*Path2d) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$_.arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

// ArcTo fn
// js:"arcTo"
func (*Path2d) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	js.Rewrite("$_.arcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

// BezierCurveTo fn
// js:"bezierCurveTo"
func (*Path2d) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	js.Rewrite("$_.bezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

// ClosePath fn
// js:"closePath"
func (*Path2d) ClosePath() {
	js.Rewrite("$_.closePath()")
}

// Ellipse fn
// js:"ellipse"
func (*Path2d) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$_.ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

// LineTo fn
// js:"lineTo"
func (*Path2d) LineTo(x float32, y float32) {
	js.Rewrite("$_.lineTo($1, $2)", x, y)
}

// MoveTo fn
// js:"moveTo"
func (*Path2d) MoveTo(x float32, y float32) {
	js.Rewrite("$_.moveTo($1, $2)", x, y)
}

// QuadraticCurveTo fn
// js:"quadraticCurveTo"
func (*Path2d) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	js.Rewrite("$_.quadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

// Rect fn
// js:"rect"
func (*Path2d) Rect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$_.rect($1, $2, $3, $4)", x, y, w, h)
}
