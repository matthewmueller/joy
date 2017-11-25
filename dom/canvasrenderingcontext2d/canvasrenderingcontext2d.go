package canvasrenderingcontext2d

import (
	"github.com/matthewmueller/golly/dom/canvasfillrule"
	"github.com/matthewmueller/golly/dom/canvasgradient"
	"github.com/matthewmueller/golly/dom/canvaspattern"
	"github.com/matthewmueller/golly/dom/htmlcanvaselement"
	"github.com/matthewmueller/golly/dom/imagedata"
	"github.com/matthewmueller/golly/dom/path2d"
	"github.com/matthewmueller/golly/dom/textmetrics"
	"github.com/matthewmueller/golly/dom/window"
	"github.com/matthewmueller/golly/js"
)

// CanvasRenderingContext2d struct
// js:"CanvasRenderingContext2d,omit"
type CanvasRenderingContext2d struct {
}

// Arc fn
func (*CanvasRenderingContext2d) Arc(x float32, y float32, radius float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$<.arc($1, $2, $3, $4, $5, $6)", x, y, radius, startAngle, endAngle, anticlockwise)
}

// ArcTo fn
func (*CanvasRenderingContext2d) ArcTo(x1 float32, y1 float32, x2 float32, y2 float32, radius float32) {
	js.Rewrite("$<.arcTo($1, $2, $3, $4, $5)", x1, y1, x2, y2, radius)
}

// BezierCurveTo fn
func (*CanvasRenderingContext2d) BezierCurveTo(cp1x float32, cp1y float32, cp2x float32, cp2y float32, x float32, y float32) {
	js.Rewrite("$<.bezierCurveTo($1, $2, $3, $4, $5, $6)", cp1x, cp1y, cp2x, cp2y, x, y)
}

// ClosePath fn
func (*CanvasRenderingContext2d) ClosePath() {
	js.Rewrite("$<.closePath()")
}

// Ellipse fn
func (*CanvasRenderingContext2d) Ellipse(x float32, y float32, radiusX float32, radiusY float32, rotation float32, startAngle float32, endAngle float32, anticlockwise *bool) {
	js.Rewrite("$<.ellipse($1, $2, $3, $4, $5, $6, $7, $8)", x, y, radiusX, radiusY, rotation, startAngle, endAngle, anticlockwise)
}

// LineTo fn
func (*CanvasRenderingContext2d) LineTo(x float32, y float32) {
	js.Rewrite("$<.lineTo($1, $2)", x, y)
}

// MoveTo fn
func (*CanvasRenderingContext2d) MoveTo(x float32, y float32) {
	js.Rewrite("$<.moveTo($1, $2)", x, y)
}

// QuadraticCurveTo fn
func (*CanvasRenderingContext2d) QuadraticCurveTo(cpx float32, cpy float32, x float32, y float32) {
	js.Rewrite("$<.quadraticCurveTo($1, $2, $3, $4)", cpx, cpy, x, y)
}

// Rect fn
func (*CanvasRenderingContext2d) Rect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.rect($1, $2, $3, $4)", x, y, w, h)
}

// BeginPath fn
func (*CanvasRenderingContext2d) BeginPath() {
	js.Rewrite("$<.beginPath()")
}

// ClearRect fn
func (*CanvasRenderingContext2d) ClearRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.clearRect($1, $2, $3, $4)", x, y, w, h)
}

// Clip fn
func (*CanvasRenderingContext2d) Clip(fillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.clip($1)", fillRule)
}

// CreateImageData fn
func (*CanvasRenderingContext2d) CreateImageData(imageDataOrSw interface{}, sh *float32) (i *imagedata.ImageData) {
	js.Rewrite("$<.createImageData($1, $2)", imageDataOrSw, sh)
	return i
}

// CreateLinearGradient fn
func (*CanvasRenderingContext2d) CreateLinearGradient(x0 float32, y0 float32, x1 float32, y1 float32) (c *canvasgradient.CanvasGradient) {
	js.Rewrite("$<.createLinearGradient($1, $2, $3, $4)", x0, y0, x1, y1)
	return c
}

// CreatePattern fn
func (*CanvasRenderingContext2d) CreatePattern(image interface{}, repetition string) (c *canvaspattern.CanvasPattern) {
	js.Rewrite("$<.createPattern($1, $2)", image, repetition)
	return c
}

// CreateRadialGradient fn
func (*CanvasRenderingContext2d) CreateRadialGradient(x0 float32, y0 float32, r0 float32, x1 float32, y1 float32, r1 float32) (c *canvasgradient.CanvasGradient) {
	js.Rewrite("$<.createRadialGradient($1, $2, $3, $4, $5, $6)", x0, y0, r0, x1, y1, r1)
	return c
}

// DrawFocusIfNeeded fn
func (*CanvasRenderingContext2d) DrawFocusIfNeeded(element window.Element) {
	js.Rewrite("$<.drawFocusIfNeeded($1)", element)
}

// DrawImage fn
func (*CanvasRenderingContext2d) DrawImage(image interface{}, offsetX float32, offsetY float32, width *float32, height *float32, canvasOffsetX *float32, canvasOffsetY *float32, canvasImageWidth *float32, canvasImageHeight *float32) {
	js.Rewrite("$<.drawImage($1, $2, $3, $4, $5, $6, $7, $8, $9)", image, offsetX, offsetY, width, height, canvasOffsetX, canvasOffsetY, canvasImageWidth, canvasImageHeight)
}

// Fill fn
func (*CanvasRenderingContext2d) Fill(fillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.fill($1)", fillRule)
}

// FillRect fn
func (*CanvasRenderingContext2d) FillRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.fillRect($1, $2, $3, $4)", x, y, w, h)
}

// FillText fn
func (*CanvasRenderingContext2d) FillText(text string, x float32, y float32, maxWidth *float32) {
	js.Rewrite("$<.fillText($1, $2, $3, $4)", text, x, y, maxWidth)
}

// GetImageData fn
func (*CanvasRenderingContext2d) GetImageData(sx float32, sy float32, sw float32, sh float32) (i *imagedata.ImageData) {
	js.Rewrite("$<.getImageData($1, $2, $3, $4)", sx, sy, sw, sh)
	return i
}

// GetLineDash fn
func (*CanvasRenderingContext2d) GetLineDash() (f []float32) {
	js.Rewrite("$<.getLineDash()")
	return f
}

// IsPointInPath fn
func (*CanvasRenderingContext2d) IsPointInPath(x float32, y float32, fillRule *canvasfillrule.CanvasFillRule) (b bool) {
	js.Rewrite("$<.isPointInPath($1, $2, $3)", x, y, fillRule)
	return b
}

// MeasureText fn
func (*CanvasRenderingContext2d) MeasureText(text string) (t *textmetrics.TextMetrics) {
	js.Rewrite("$<.measureText($1)", text)
	return t
}

// PutImageData fn
func (*CanvasRenderingContext2d) PutImageData(imagedata *imagedata.ImageData, dx float32, dy float32, dirtyX *float32, dirtyY *float32, dirtyWidth *float32, dirtyHeight *float32) {
	js.Rewrite("$<.putImageData($1, $2, $3, $4, $5, $6, $7)", imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

// Restore fn
func (*CanvasRenderingContext2d) Restore() {
	js.Rewrite("$<.restore()")
}

// Rotate fn
func (*CanvasRenderingContext2d) Rotate(angle float32) {
	js.Rewrite("$<.rotate($1)", angle)
}

// Save fn
func (*CanvasRenderingContext2d) Save() {
	js.Rewrite("$<.save()")
}

// Scale fn
func (*CanvasRenderingContext2d) Scale(x float32, y float32) {
	js.Rewrite("$<.scale($1, $2)", x, y)
}

// SetLineDash fn
func (*CanvasRenderingContext2d) SetLineDash(segments []float32) {
	js.Rewrite("$<.setLineDash($1)", segments)
}

// SetTransform fn
func (*CanvasRenderingContext2d) SetTransform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.setTransform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

// Stroke fn
func (*CanvasRenderingContext2d) Stroke(path *path2d.Path2d) {
	js.Rewrite("$<.stroke($1)", path)
}

// StrokeRect fn
func (*CanvasRenderingContext2d) StrokeRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.strokeRect($1, $2, $3, $4)", x, y, w, h)
}

// StrokeText fn
func (*CanvasRenderingContext2d) StrokeText(text string, x float32, y float32, maxWidth *float32) {
	js.Rewrite("$<.strokeText($1, $2, $3, $4)", text, x, y, maxWidth)
}

// Transform fn
func (*CanvasRenderingContext2d) Transform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.transform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

// Translate fn
func (*CanvasRenderingContext2d) Translate(x float32, y float32) {
	js.Rewrite("$<.translate($1, $2)", x, y)
}

// Canvas prop
func (*CanvasRenderingContext2D) Canvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	js.Rewrite("$<.canvas")
	return canvas
}

// FillStyle prop
func (*CanvasRenderingContext2D) FillStyle() (fillStyle interface{}) {
	js.Rewrite("$<.fillStyle")
	return fillStyle
}

// FillStyle prop
func (*CanvasRenderingContext2D) SetFillStyle(fillStyle interface{}) {
	js.Rewrite("$<.fillStyle = $1", fillStyle)
}

// Font prop
func (*CanvasRenderingContext2D) Font() (font string) {
	js.Rewrite("$<.font")
	return font
}

// Font prop
func (*CanvasRenderingContext2D) SetFont(font string) {
	js.Rewrite("$<.font = $1", font)
}

// GlobalAlpha prop
func (*CanvasRenderingContext2D) GlobalAlpha() (globalAlpha float32) {
	js.Rewrite("$<.globalAlpha")
	return globalAlpha
}

// GlobalAlpha prop
func (*CanvasRenderingContext2D) SetGlobalAlpha(globalAlpha float32) {
	js.Rewrite("$<.globalAlpha = $1", globalAlpha)
}

// GlobalCompositeOperation prop
func (*CanvasRenderingContext2D) GlobalCompositeOperation() (globalCompositeOperation string) {
	js.Rewrite("$<.globalCompositeOperation")
	return globalCompositeOperation
}

// GlobalCompositeOperation prop
func (*CanvasRenderingContext2D) SetGlobalCompositeOperation(globalCompositeOperation string) {
	js.Rewrite("$<.globalCompositeOperation = $1", globalCompositeOperation)
}

// ImageSmoothingEnabled prop
func (*CanvasRenderingContext2D) ImageSmoothingEnabled() (imageSmoothingEnabled bool) {
	js.Rewrite("$<.imageSmoothingEnabled")
	return imageSmoothingEnabled
}

// ImageSmoothingEnabled prop
func (*CanvasRenderingContext2D) SetImageSmoothingEnabled(imageSmoothingEnabled bool) {
	js.Rewrite("$<.imageSmoothingEnabled = $1", imageSmoothingEnabled)
}

// LineCap prop
func (*CanvasRenderingContext2D) LineCap() (lineCap string) {
	js.Rewrite("$<.lineCap")
	return lineCap
}

// LineCap prop
func (*CanvasRenderingContext2D) SetLineCap(lineCap string) {
	js.Rewrite("$<.lineCap = $1", lineCap)
}

// LineDashOffset prop
func (*CanvasRenderingContext2D) LineDashOffset() (lineDashOffset float32) {
	js.Rewrite("$<.lineDashOffset")
	return lineDashOffset
}

// LineDashOffset prop
func (*CanvasRenderingContext2D) SetLineDashOffset(lineDashOffset float32) {
	js.Rewrite("$<.lineDashOffset = $1", lineDashOffset)
}

// LineJoin prop
func (*CanvasRenderingContext2D) LineJoin() (lineJoin string) {
	js.Rewrite("$<.lineJoin")
	return lineJoin
}

// LineJoin prop
func (*CanvasRenderingContext2D) SetLineJoin(lineJoin string) {
	js.Rewrite("$<.lineJoin = $1", lineJoin)
}

// LineWidth prop
func (*CanvasRenderingContext2D) LineWidth() (lineWidth float32) {
	js.Rewrite("$<.lineWidth")
	return lineWidth
}

// LineWidth prop
func (*CanvasRenderingContext2D) SetLineWidth(lineWidth float32) {
	js.Rewrite("$<.lineWidth = $1", lineWidth)
}

// MiterLimit prop
func (*CanvasRenderingContext2D) MiterLimit() (miterLimit float32) {
	js.Rewrite("$<.miterLimit")
	return miterLimit
}

// MiterLimit prop
func (*CanvasRenderingContext2D) SetMiterLimit(miterLimit float32) {
	js.Rewrite("$<.miterLimit = $1", miterLimit)
}

// MsFillRule prop
func (*CanvasRenderingContext2D) MsFillRule() (msFillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.msFillRule")
	return msFillRule
}

// MsFillRule prop
func (*CanvasRenderingContext2D) SetMsFillRule(msFillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.msFillRule = $1", msFillRule)
}

// ShadowBlur prop
func (*CanvasRenderingContext2D) ShadowBlur() (shadowBlur float32) {
	js.Rewrite("$<.shadowBlur")
	return shadowBlur
}

// ShadowBlur prop
func (*CanvasRenderingContext2D) SetShadowBlur(shadowBlur float32) {
	js.Rewrite("$<.shadowBlur = $1", shadowBlur)
}

// ShadowColor prop
func (*CanvasRenderingContext2D) ShadowColor() (shadowColor string) {
	js.Rewrite("$<.shadowColor")
	return shadowColor
}

// ShadowColor prop
func (*CanvasRenderingContext2D) SetShadowColor(shadowColor string) {
	js.Rewrite("$<.shadowColor = $1", shadowColor)
}

// ShadowOffsetX prop
func (*CanvasRenderingContext2D) ShadowOffsetX() (shadowOffsetX float32) {
	js.Rewrite("$<.shadowOffsetX")
	return shadowOffsetX
}

// ShadowOffsetX prop
func (*CanvasRenderingContext2D) SetShadowOffsetX(shadowOffsetX float32) {
	js.Rewrite("$<.shadowOffsetX = $1", shadowOffsetX)
}

// ShadowOffsetY prop
func (*CanvasRenderingContext2D) ShadowOffsetY() (shadowOffsetY float32) {
	js.Rewrite("$<.shadowOffsetY")
	return shadowOffsetY
}

// ShadowOffsetY prop
func (*CanvasRenderingContext2D) SetShadowOffsetY(shadowOffsetY float32) {
	js.Rewrite("$<.shadowOffsetY = $1", shadowOffsetY)
}

// StrokeStyle prop
func (*CanvasRenderingContext2D) StrokeStyle() (strokeStyle interface{}) {
	js.Rewrite("$<.strokeStyle")
	return strokeStyle
}

// StrokeStyle prop
func (*CanvasRenderingContext2D) SetStrokeStyle(strokeStyle interface{}) {
	js.Rewrite("$<.strokeStyle = $1", strokeStyle)
}

// TextAlign prop
func (*CanvasRenderingContext2D) TextAlign() (textAlign string) {
	js.Rewrite("$<.textAlign")
	return textAlign
}

// TextAlign prop
func (*CanvasRenderingContext2D) SetTextAlign(textAlign string) {
	js.Rewrite("$<.textAlign = $1", textAlign)
}

// TextBaseline prop
func (*CanvasRenderingContext2D) TextBaseline() (textBaseline string) {
	js.Rewrite("$<.textBaseline")
	return textBaseline
}

// TextBaseline prop
func (*CanvasRenderingContext2D) SetTextBaseline(textBaseline string) {
	js.Rewrite("$<.textBaseline = $1", textBaseline)
}
