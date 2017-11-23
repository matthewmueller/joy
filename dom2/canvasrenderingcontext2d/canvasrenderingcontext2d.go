package canvasrenderingcontext2d

import (
	"github.com/matthewmueller/golly/dom2/canvasfillrule"
	"github.com/matthewmueller/golly/dom2/canvasgradient"
	"github.com/matthewmueller/golly/dom2/canvaspattern"
	"github.com/matthewmueller/golly/dom2/htmlcanvaselement"
	"github.com/matthewmueller/golly/dom2/imagedata"
	"github.com/matthewmueller/golly/dom2/path2d"
	"github.com/matthewmueller/golly/dom2/textmetrics"
	"github.com/matthewmueller/golly/dom2/window"
	"github.com/matthewmueller/golly/js"
)

// CanvasRenderingContext2d struct
// js:"CanvasRenderingContext2d,omit"
type CanvasRenderingContext2d struct {
}

// Arc
func (*CanvasRenderingContext2D) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$<.Arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

// ArcTo
func (*CanvasRenderingContext2D) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	js.Rewrite("$<.ArcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

// BezierCurveTo
func (*CanvasRenderingContext2D) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	js.Rewrite("$<.BezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

// ClosePath
func (*CanvasRenderingContext2D) ClosePath() {
	js.Rewrite("$<.ClosePath()")
}

// Ellipse
func (*CanvasRenderingContext2D) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$<.Ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

// LineTo
func (*CanvasRenderingContext2D) LineTo(x float32, y float32) {
	js.Rewrite("$<.LineTo($1, $2)", x, y)
}

// MoveTo
func (*CanvasRenderingContext2D) MoveTo(x float32, y float32) {
	js.Rewrite("$<.MoveTo($1, $2)", x, y)
}

// QuadraticCurveTo
func (*CanvasRenderingContext2D) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	js.Rewrite("$<.QuadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

// Rect
func (*CanvasRenderingContext2D) Rect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.Rect($1, $2, $3, $4)", x, y, w, h)
}

// BeginPath
func (*CanvasRenderingContext2D) BeginPath() {
	js.Rewrite("$<.BeginPath()")
}

// ClearRect
func (*CanvasRenderingContext2D) ClearRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.ClearRect($1, $2, $3, $4)", x, y, w, h)
}

// Clip
func (*CanvasRenderingContext2D) Clip(fillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.Clip($1)", fillRule)
}

// CreateImageData
func (*CanvasRenderingContext2D) CreateImageData(imageDataOrSw interface{}, sh *float32) (i *imagedata.ImageData) {
	js.Rewrite("$<.CreateImageData($1, $2)", imageDataOrSw, sh)
	return i
}

// CreateLinearGradient
func (*CanvasRenderingContext2D) CreateLinearGradient(x0 float32, y0 float32, x1 float32, y1 float32) (c *canvasgradient.CanvasGradient) {
	js.Rewrite("$<.CreateLinearGradient($1, $2, $3, $4)", x0, y0, x1, y1)
	return c
}

// CreatePattern
func (*CanvasRenderingContext2D) CreatePattern(image interface{}, repetition string) (c *canvaspattern.CanvasPattern) {
	js.Rewrite("$<.CreatePattern($1, $2)", image, repetition)
	return c
}

// CreateRadialGradient
func (*CanvasRenderingContext2D) CreateRadialGradient(x0 float32, y0 float32, r0 float32, x1 float32, y1 float32, r1 float32) (c *canvasgradient.CanvasGradient) {
	js.Rewrite("$<.CreateRadialGradient($1, $2, $3, $4, $5, $6)", x0, y0, r0, x1, y1, r1)
	return c
}

// DrawFocusIfNeeded
func (*CanvasRenderingContext2D) DrawFocusIfNeeded(element window.Element) {
	js.Rewrite("$<.DrawFocusIfNeeded($1)", element)
}

// DrawImage
func (*CanvasRenderingContext2D) DrawImage(image interface{}, offsetX float32, offsetY float32, width *float32, height *float32, canvasOffsetX *float32, canvasOffsetY *float32, canvasImageWidth *float32, canvasImageHeight *float32) {
	js.Rewrite("$<.DrawImage($1, $2, $3, $4, $5, $6, $7, $8, $9)", image, offsetX, offsetY, width, height, canvasOffsetX, canvasOffsetY, canvasImageWidth, canvasImageHeight)
}

// Fill
func (*CanvasRenderingContext2D) Fill(fillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.Fill($1)", fillRule)
}

// FillRect
func (*CanvasRenderingContext2D) FillRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.FillRect($1, $2, $3, $4)", x, y, w, h)
}

// FillText
func (*CanvasRenderingContext2D) FillText(text string, x float32, y float32, maxWidth *float32) {
	js.Rewrite("$<.FillText($1, $2, $3, $4)", text, x, y, maxWidth)
}

// GetImageData
func (*CanvasRenderingContext2D) GetImageData(sx float32, sy float32, sw float32, sh float32) (i *imagedata.ImageData) {
	js.Rewrite("$<.GetImageData($1, $2, $3, $4)", sx, sy, sw, sh)
	return i
}

// GetLineDash
func (*CanvasRenderingContext2D) GetLineDash() (f []float32) {
	js.Rewrite("$<.GetLineDash()")
	return f
}

// IsPointInPath
func (*CanvasRenderingContext2D) IsPointInPath(x float32, y float32, fillRule *canvasfillrule.CanvasFillRule) (b bool) {
	js.Rewrite("$<.IsPointInPath($1, $2, $3)", x, y, fillRule)
	return b
}

// MeasureText
func (*CanvasRenderingContext2D) MeasureText(text string) (t *textmetrics.TextMetrics) {
	js.Rewrite("$<.MeasureText($1)", text)
	return t
}

// PutImageData
func (*CanvasRenderingContext2D) PutImageData(imagedata *imagedata.ImageData, dx float32, dy float32, dirtyX *float32, dirtyY *float32, dirtyWidth *float32, dirtyHeight *float32) {
	js.Rewrite("$<.PutImageData($1, $2, $3, $4, $5, $6, $7)", imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

// Restore
func (*CanvasRenderingContext2D) Restore() {
	js.Rewrite("$<.Restore()")
}

// Rotate
func (*CanvasRenderingContext2D) Rotate(angle float32) {
	js.Rewrite("$<.Rotate($1)", angle)
}

// Save
func (*CanvasRenderingContext2D) Save() {
	js.Rewrite("$<.Save()")
}

// Scale
func (*CanvasRenderingContext2D) Scale(x float32, y float32) {
	js.Rewrite("$<.Scale($1, $2)", x, y)
}

// SetLineDash
func (*CanvasRenderingContext2D) SetLineDash(segments []float32) {
	js.Rewrite("$<.SetLineDash($1)", segments)
}

// SetTransform
func (*CanvasRenderingContext2D) SetTransform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.SetTransform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

// Stroke
func (*CanvasRenderingContext2D) Stroke(path *path2d.Path2d) {
	js.Rewrite("$<.Stroke($1)", path)
}

// StrokeRect
func (*CanvasRenderingContext2D) StrokeRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.StrokeRect($1, $2, $3, $4)", x, y, w, h)
}

// StrokeText
func (*CanvasRenderingContext2D) StrokeText(text string, x float32, y float32, maxWidth *float32) {
	js.Rewrite("$<.StrokeText($1, $2, $3, $4)", text, x, y, maxWidth)
}

// Transform
func (*CanvasRenderingContext2D) Transform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.Transform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

// Translate
func (*CanvasRenderingContext2D) Translate(x float32, y float32) {
	js.Rewrite("$<.Translate($1, $2)", x, y)
}

// Canvas
func (*CanvasRenderingContext2D) Canvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	js.Rewrite("$<.Canvas")
	return canvas
}

// FillStyle
func (*CanvasRenderingContext2D) FillStyle() (fillStyle interface{}) {
	js.Rewrite("$<.FillStyle")
	return fillStyle
}

// FillStyle
func (*CanvasRenderingContext2D) SetFillStyle(fillStyle interface{}) {
	js.Rewrite("$<.FillStyle = $1", fillStyle)
}

// Font
func (*CanvasRenderingContext2D) Font() (font string) {
	js.Rewrite("$<.Font")
	return font
}

// Font
func (*CanvasRenderingContext2D) SetFont(font string) {
	js.Rewrite("$<.Font = $1", font)
}

// GlobalAlpha
func (*CanvasRenderingContext2D) GlobalAlpha() (globalAlpha float32) {
	js.Rewrite("$<.GlobalAlpha")
	return globalAlpha
}

// GlobalAlpha
func (*CanvasRenderingContext2D) SetGlobalAlpha(globalAlpha float32) {
	js.Rewrite("$<.GlobalAlpha = $1", globalAlpha)
}

// GlobalCompositeOperation
func (*CanvasRenderingContext2D) GlobalCompositeOperation() (globalCompositeOperation string) {
	js.Rewrite("$<.GlobalCompositeOperation")
	return globalCompositeOperation
}

// GlobalCompositeOperation
func (*CanvasRenderingContext2D) SetGlobalCompositeOperation(globalCompositeOperation string) {
	js.Rewrite("$<.GlobalCompositeOperation = $1", globalCompositeOperation)
}

// ImageSmoothingEnabled
func (*CanvasRenderingContext2D) ImageSmoothingEnabled() (imageSmoothingEnabled bool) {
	js.Rewrite("$<.ImageSmoothingEnabled")
	return imageSmoothingEnabled
}

// ImageSmoothingEnabled
func (*CanvasRenderingContext2D) SetImageSmoothingEnabled(imageSmoothingEnabled bool) {
	js.Rewrite("$<.ImageSmoothingEnabled = $1", imageSmoothingEnabled)
}

// LineCap
func (*CanvasRenderingContext2D) LineCap() (lineCap string) {
	js.Rewrite("$<.LineCap")
	return lineCap
}

// LineCap
func (*CanvasRenderingContext2D) SetLineCap(lineCap string) {
	js.Rewrite("$<.LineCap = $1", lineCap)
}

// LineDashOffset
func (*CanvasRenderingContext2D) LineDashOffset() (lineDashOffset float32) {
	js.Rewrite("$<.LineDashOffset")
	return lineDashOffset
}

// LineDashOffset
func (*CanvasRenderingContext2D) SetLineDashOffset(lineDashOffset float32) {
	js.Rewrite("$<.LineDashOffset = $1", lineDashOffset)
}

// LineJoin
func (*CanvasRenderingContext2D) LineJoin() (lineJoin string) {
	js.Rewrite("$<.LineJoin")
	return lineJoin
}

// LineJoin
func (*CanvasRenderingContext2D) SetLineJoin(lineJoin string) {
	js.Rewrite("$<.LineJoin = $1", lineJoin)
}

// LineWidth
func (*CanvasRenderingContext2D) LineWidth() (lineWidth float32) {
	js.Rewrite("$<.LineWidth")
	return lineWidth
}

// LineWidth
func (*CanvasRenderingContext2D) SetLineWidth(lineWidth float32) {
	js.Rewrite("$<.LineWidth = $1", lineWidth)
}

// MiterLimit
func (*CanvasRenderingContext2D) MiterLimit() (miterLimit float32) {
	js.Rewrite("$<.MiterLimit")
	return miterLimit
}

// MiterLimit
func (*CanvasRenderingContext2D) SetMiterLimit(miterLimit float32) {
	js.Rewrite("$<.MiterLimit = $1", miterLimit)
}

// MsFillRule
func (*CanvasRenderingContext2D) MsFillRule() (msFillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.MsFillRule")
	return msFillRule
}

// MsFillRule
func (*CanvasRenderingContext2D) SetMsFillRule(msFillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.MsFillRule = $1", msFillRule)
}

// ShadowBlur
func (*CanvasRenderingContext2D) ShadowBlur() (shadowBlur float32) {
	js.Rewrite("$<.ShadowBlur")
	return shadowBlur
}

// ShadowBlur
func (*CanvasRenderingContext2D) SetShadowBlur(shadowBlur float32) {
	js.Rewrite("$<.ShadowBlur = $1", shadowBlur)
}

// ShadowColor
func (*CanvasRenderingContext2D) ShadowColor() (shadowColor string) {
	js.Rewrite("$<.ShadowColor")
	return shadowColor
}

// ShadowColor
func (*CanvasRenderingContext2D) SetShadowColor(shadowColor string) {
	js.Rewrite("$<.ShadowColor = $1", shadowColor)
}

// ShadowOffsetX
func (*CanvasRenderingContext2D) ShadowOffsetX() (shadowOffsetX float32) {
	js.Rewrite("$<.ShadowOffsetX")
	return shadowOffsetX
}

// ShadowOffsetX
func (*CanvasRenderingContext2D) SetShadowOffsetX(shadowOffsetX float32) {
	js.Rewrite("$<.ShadowOffsetX = $1", shadowOffsetX)
}

// ShadowOffsetY
func (*CanvasRenderingContext2D) ShadowOffsetY() (shadowOffsetY float32) {
	js.Rewrite("$<.ShadowOffsetY")
	return shadowOffsetY
}

// ShadowOffsetY
func (*CanvasRenderingContext2D) SetShadowOffsetY(shadowOffsetY float32) {
	js.Rewrite("$<.ShadowOffsetY = $1", shadowOffsetY)
}

// StrokeStyle
func (*CanvasRenderingContext2D) StrokeStyle() (strokeStyle interface{}) {
	js.Rewrite("$<.StrokeStyle")
	return strokeStyle
}

// StrokeStyle
func (*CanvasRenderingContext2D) SetStrokeStyle(strokeStyle interface{}) {
	js.Rewrite("$<.StrokeStyle = $1", strokeStyle)
}

// TextAlign
func (*CanvasRenderingContext2D) TextAlign() (textAlign string) {
	js.Rewrite("$<.TextAlign")
	return textAlign
}

// TextAlign
func (*CanvasRenderingContext2D) SetTextAlign(textAlign string) {
	js.Rewrite("$<.TextAlign = $1", textAlign)
}

// TextBaseline
func (*CanvasRenderingContext2D) TextBaseline() (textBaseline string) {
	js.Rewrite("$<.TextBaseline")
	return textBaseline
}

// TextBaseline
func (*CanvasRenderingContext2D) SetTextBaseline(textBaseline string) {
	js.Rewrite("$<.TextBaseline = $1", textBaseline)
}
