package canvasrenderingcontext2d

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/canvasfillrule"
	"github.com/matthewmueller/golly/internal/gendom/dom/canvasgradient"
	"github.com/matthewmueller/golly/internal/gendom/dom/canvaspathmethods"
	"github.com/matthewmueller/golly/internal/gendom/dom/canvaspattern"
	"github.com/matthewmueller/golly/internal/gendom/dom/element"
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcanvaselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/imagedata"
	"github.com/matthewmueller/golly/internal/gendom/dom/path2d"
	"github.com/matthewmueller/golly/internal/gendom/dom/textmetrics"
	"github.com/matthewmueller/golly/js"
)

type CanvasRenderingContext2D struct {
	*canvaspathmethods.CanvasPathMethods
}

func (*CanvasRenderingContext2D) BeginPath() {
	js.Rewrite("$<.beginPath()")
}

func (*CanvasRenderingContext2D) ClearRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.clearRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2D) Clip(fillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.clip($1)", fillRule)
}

func (*CanvasRenderingContext2D) CreateImageData(imageDataOrSw interface{}, sh float32) (i *imagedata.ImageData) {
	js.Rewrite("$<.createImageData($1, $2)", imageDataOrSw, sh)
	return i
}

func (*CanvasRenderingContext2D) CreateLinearGradient(x0 float32, y0 float32, x1 float32, y1 float32) (c *canvasgradient.CanvasGradient) {
	js.Rewrite("$<.createLinearGradient($1, $2, $3, $4)", x0, y0, x1, y1)
	return c
}

func (*CanvasRenderingContext2D) CreatePattern(image interface{}, repetition string) (c *canvaspattern.CanvasPattern) {
	js.Rewrite("$<.createPattern($1, $2)", image, repetition)
	return c
}

func (*CanvasRenderingContext2D) CreateRadialGradient(x0 float32, y0 float32, r0 float32, x1 float32, y1 float32, r1 float32) (c *canvasgradient.CanvasGradient) {
	js.Rewrite("$<.createRadialGradient($1, $2, $3, $4, $5, $6)", x0, y0, r0, x1, y1, r1)
	return c
}

func (*CanvasRenderingContext2D) DrawFocusIfNeeded(element *element.Element) {
	js.Rewrite("$<.drawFocusIfNeeded($1)", element)
}

func (*CanvasRenderingContext2D) DrawImage(image interface{}, offsetX float32, offsetY float32, width float32, height float32, canvasOffsetX float32, canvasOffsetY float32, canvasImageWidth float32, canvasImageHeight float32) {
	js.Rewrite("$<.drawImage($1, $2, $3, $4, $5, $6, $7, $8, $9)", image, offsetX, offsetY, width, height, canvasOffsetX, canvasOffsetY, canvasImageWidth, canvasImageHeight)
}

func (*CanvasRenderingContext2D) Fill(fillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.fill($1)", fillRule)
}

func (*CanvasRenderingContext2D) FillRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.fillRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2D) FillText(text string, x float32, y float32, maxWidth float32) {
	js.Rewrite("$<.fillText($1, $2, $3, $4)", text, x, y, maxWidth)
}

func (*CanvasRenderingContext2D) GetImageData(sx float32, sy float32, sw float32, sh float32) (i *imagedata.ImageData) {
	js.Rewrite("$<.getImageData($1, $2, $3, $4)", sx, sy, sw, sh)
	return i
}

func (*CanvasRenderingContext2D) GetLineDash() (f []float32) {
	js.Rewrite("$<.getLineDash()")
	return f
}

func (*CanvasRenderingContext2D) IsPointInPath(x float32, y float32, fillRule *canvasfillrule.CanvasFillRule) (b bool) {
	js.Rewrite("$<.isPointInPath($1, $2, $3)", x, y, fillRule)
	return b
}

func (*CanvasRenderingContext2D) MeasureText(text string) (t *textmetrics.TextMetrics) {
	js.Rewrite("$<.measureText($1)", text)
	return t
}

func (*CanvasRenderingContext2D) PutImageData(imagedata *imagedata.ImageData, dx float32, dy float32, dirtyX float32, dirtyY float32, dirtyWidth float32, dirtyHeight float32) {
	js.Rewrite("$<.putImageData($1, $2, $3, $4, $5, $6, $7)", imagedata, dx, dy, dirtyX, dirtyY, dirtyWidth, dirtyHeight)
}

func (*CanvasRenderingContext2D) Restore() {
	js.Rewrite("$<.restore()")
}

func (*CanvasRenderingContext2D) Rotate(angle float32) {
	js.Rewrite("$<.rotate($1)", angle)
}

func (*CanvasRenderingContext2D) Save() {
	js.Rewrite("$<.save()")
}

func (*CanvasRenderingContext2D) Scale(x float32, y float32) {
	js.Rewrite("$<.scale($1, $2)", x, y)
}

func (*CanvasRenderingContext2D) SetLineDash(segments []float32) {
	js.Rewrite("$<.setLineDash($1)", segments)
}

func (*CanvasRenderingContext2D) SetTransform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.setTransform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

func (*CanvasRenderingContext2D) Stroke(path *path2d.Path2D) {
	js.Rewrite("$<.stroke($1)", path)
}

func (*CanvasRenderingContext2D) StrokeRect(x float32, y float32, w float32, h float32) {
	js.Rewrite("$<.strokeRect($1, $2, $3, $4)", x, y, w, h)
}

func (*CanvasRenderingContext2D) StrokeText(text string, x float32, y float32, maxWidth float32) {
	js.Rewrite("$<.strokeText($1, $2, $3, $4)", text, x, y, maxWidth)
}

func (*CanvasRenderingContext2D) Transform(m11 float32, m12 float32, m21 float32, m22 float32, dx float32, dy float32) {
	js.Rewrite("$<.transform($1, $2, $3, $4, $5, $6)", m11, m12, m21, m22, dx, dy)
}

func (*CanvasRenderingContext2D) Translate(x float32, y float32) {
	js.Rewrite("$<.translate($1, $2)", x, y)
}

func (*CanvasRenderingContext2D) GetCanvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	js.Rewrite("$<.canvas")
	return canvas
}

func (*CanvasRenderingContext2D) GetFillStyle() (fillStyle interface{}) {
	js.Rewrite("$<.fillStyle")
	return fillStyle
}

func (*CanvasRenderingContext2D) SetFillStyle(fillStyle interface{}) {
	js.Rewrite("$<.fillStyle = $1", fillStyle)
}

func (*CanvasRenderingContext2D) GetFont() (font string) {
	js.Rewrite("$<.font")
	return font
}

func (*CanvasRenderingContext2D) SetFont(font string) {
	js.Rewrite("$<.font = $1", font)
}

func (*CanvasRenderingContext2D) GetGlobalAlpha() (globalAlpha float32) {
	js.Rewrite("$<.globalAlpha")
	return globalAlpha
}

func (*CanvasRenderingContext2D) SetGlobalAlpha(globalAlpha float32) {
	js.Rewrite("$<.globalAlpha = $1", globalAlpha)
}

func (*CanvasRenderingContext2D) GetGlobalCompositeOperation() (globalCompositeOperation string) {
	js.Rewrite("$<.globalCompositeOperation")
	return globalCompositeOperation
}

func (*CanvasRenderingContext2D) SetGlobalCompositeOperation(globalCompositeOperation string) {
	js.Rewrite("$<.globalCompositeOperation = $1", globalCompositeOperation)
}

func (*CanvasRenderingContext2D) GetImageSmoothingEnabled() (imageSmoothingEnabled bool) {
	js.Rewrite("$<.imageSmoothingEnabled")
	return imageSmoothingEnabled
}

func (*CanvasRenderingContext2D) SetImageSmoothingEnabled(imageSmoothingEnabled bool) {
	js.Rewrite("$<.imageSmoothingEnabled = $1", imageSmoothingEnabled)
}

func (*CanvasRenderingContext2D) GetLineCap() (lineCap string) {
	js.Rewrite("$<.lineCap")
	return lineCap
}

func (*CanvasRenderingContext2D) SetLineCap(lineCap string) {
	js.Rewrite("$<.lineCap = $1", lineCap)
}

func (*CanvasRenderingContext2D) GetLineDashOffset() (lineDashOffset float32) {
	js.Rewrite("$<.lineDashOffset")
	return lineDashOffset
}

func (*CanvasRenderingContext2D) SetLineDashOffset(lineDashOffset float32) {
	js.Rewrite("$<.lineDashOffset = $1", lineDashOffset)
}

func (*CanvasRenderingContext2D) GetLineJoin() (lineJoin string) {
	js.Rewrite("$<.lineJoin")
	return lineJoin
}

func (*CanvasRenderingContext2D) SetLineJoin(lineJoin string) {
	js.Rewrite("$<.lineJoin = $1", lineJoin)
}

func (*CanvasRenderingContext2D) GetLineWidth() (lineWidth float32) {
	js.Rewrite("$<.lineWidth")
	return lineWidth
}

func (*CanvasRenderingContext2D) SetLineWidth(lineWidth float32) {
	js.Rewrite("$<.lineWidth = $1", lineWidth)
}

func (*CanvasRenderingContext2D) GetMiterLimit() (miterLimit float32) {
	js.Rewrite("$<.miterLimit")
	return miterLimit
}

func (*CanvasRenderingContext2D) SetMiterLimit(miterLimit float32) {
	js.Rewrite("$<.miterLimit = $1", miterLimit)
}

func (*CanvasRenderingContext2D) GetMsFillRule() (msFillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.msFillRule")
	return msFillRule
}

func (*CanvasRenderingContext2D) SetMsFillRule(msFillRule *canvasfillrule.CanvasFillRule) {
	js.Rewrite("$<.msFillRule = $1", msFillRule)
}

func (*CanvasRenderingContext2D) GetShadowBlur() (shadowBlur float32) {
	js.Rewrite("$<.shadowBlur")
	return shadowBlur
}

func (*CanvasRenderingContext2D) SetShadowBlur(shadowBlur float32) {
	js.Rewrite("$<.shadowBlur = $1", shadowBlur)
}

func (*CanvasRenderingContext2D) GetShadowColor() (shadowColor string) {
	js.Rewrite("$<.shadowColor")
	return shadowColor
}

func (*CanvasRenderingContext2D) SetShadowColor(shadowColor string) {
	js.Rewrite("$<.shadowColor = $1", shadowColor)
}

func (*CanvasRenderingContext2D) GetShadowOffsetX() (shadowOffsetX float32) {
	js.Rewrite("$<.shadowOffsetX")
	return shadowOffsetX
}

func (*CanvasRenderingContext2D) SetShadowOffsetX(shadowOffsetX float32) {
	js.Rewrite("$<.shadowOffsetX = $1", shadowOffsetX)
}

func (*CanvasRenderingContext2D) GetShadowOffsetY() (shadowOffsetY float32) {
	js.Rewrite("$<.shadowOffsetY")
	return shadowOffsetY
}

func (*CanvasRenderingContext2D) SetShadowOffsetY(shadowOffsetY float32) {
	js.Rewrite("$<.shadowOffsetY = $1", shadowOffsetY)
}

func (*CanvasRenderingContext2D) GetStrokeStyle() (strokeStyle interface{}) {
	js.Rewrite("$<.strokeStyle")
	return strokeStyle
}

func (*CanvasRenderingContext2D) SetStrokeStyle(strokeStyle interface{}) {
	js.Rewrite("$<.strokeStyle = $1", strokeStyle)
}

func (*CanvasRenderingContext2D) GetTextAlign() (textAlign string) {
	js.Rewrite("$<.textAlign")
	return textAlign
}

func (*CanvasRenderingContext2D) SetTextAlign(textAlign string) {
	js.Rewrite("$<.textAlign = $1", textAlign)
}

func (*CanvasRenderingContext2D) GetTextBaseline() (textBaseline string) {
	js.Rewrite("$<.textBaseline")
	return textBaseline
}

func (*CanvasRenderingContext2D) SetTextBaseline(textBaseline string) {
	js.Rewrite("$<.textBaseline = $1", textBaseline)
}
