package canvaspathmethods

// CanvasPathMethods interface
// js:"CanvasPathMethods"
type CanvasPathMethods interface {

	// Arc
	// js:"arc"
	// jsrewrite:"$_.arc($1, $2, $3, $4, $5, $6)"
	Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool)

	// ArcTo
	// js:"arcTo"
	// jsrewrite:"$_.arcTo($1, $2, $3, $4, $5)"
	ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32)

	// BezierCurveTo
	// js:"bezierCurveTo"
	// jsrewrite:"$_.bezierCurveTo($1, $2, $3, $4, $5, $6)"
	BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32)

	// ClosePath
	// js:"closePath"
	// jsrewrite:"$_.closePath()"
	ClosePath()

	// Ellipse
	// js:"ellipse"
	// jsrewrite:"$_.ellipse($1, $2, $3, $4, $5, $6, $7, $8)"
	Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool)

	// LineTo
	// js:"lineTo"
	// jsrewrite:"$_.lineTo($1, $2)"
	LineTo(x float32, y float32)

	// MoveTo
	// js:"moveTo"
	// jsrewrite:"$_.moveTo($1, $2)"
	MoveTo(x float32, y float32)

	// QuadraticCurveTo
	// js:"quadraticCurveTo"
	// jsrewrite:"$_.quadraticCurveTo($1, $2, $3, $4)"
	QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32)

	// Rect
	// js:"rect"
	// jsrewrite:"$_.rect($1, $2, $3, $4)"
	Rect(x float32, y float32, w float32, h float32)
}
