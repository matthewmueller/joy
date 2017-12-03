package webglrenderingcontext

import (
	"github.com/matthewmueller/joy/dom/htmlcanvaselement"
	"github.com/matthewmueller/joy/dom/imagedata"
	"github.com/matthewmueller/joy/dom/webglactiveinfo"
	"github.com/matthewmueller/joy/dom/webglbuffer"
	"github.com/matthewmueller/joy/dom/webglcontextattributes"
	"github.com/matthewmueller/joy/dom/webglframebuffer"
	"github.com/matthewmueller/joy/dom/webglprogram"
	"github.com/matthewmueller/joy/dom/webglrenderbuffer"
	"github.com/matthewmueller/joy/dom/webglshader"
	"github.com/matthewmueller/joy/dom/webglshaderprecisionformat"
	"github.com/matthewmueller/joy/dom/webgltexture"
	"github.com/matthewmueller/joy/dom/webgluniformlocation"
	"github.com/matthewmueller/joy/macro"
)

// WebGLRenderingContext struct
// js:"WebGLRenderingContext,omit"
type WebGLRenderingContext struct {
}

// ActiveTexture fn
// js:"activeTexture"
func (*WebGLRenderingContext) ActiveTexture(texture uint) {
	macro.Rewrite("$_.activeTexture($1)", texture)
}

// AttachShader fn
// js:"attachShader"
func (*WebGLRenderingContext) AttachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	macro.Rewrite("$_.attachShader($1, $2)", program, shader)
}

// BindAttribLocation fn
// js:"bindAttribLocation"
func (*WebGLRenderingContext) BindAttribLocation(program *webglprogram.WebGLProgram, index uint, name string) {
	macro.Rewrite("$_.bindAttribLocation($1, $2, $3)", program, index, name)
}

// BindBuffer fn
// js:"bindBuffer"
func (*WebGLRenderingContext) BindBuffer(target uint, buffer *webglbuffer.WebGLBuffer) {
	macro.Rewrite("$_.bindBuffer($1, $2)", target, buffer)
}

// BindFramebuffer fn
// js:"bindFramebuffer"
func (*WebGLRenderingContext) BindFramebuffer(target uint, framebuffer *webglframebuffer.WebGLFramebuffer) {
	macro.Rewrite("$_.bindFramebuffer($1, $2)", target, framebuffer)
}

// BindRenderbuffer fn
// js:"bindRenderbuffer"
func (*WebGLRenderingContext) BindRenderbuffer(target uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	macro.Rewrite("$_.bindRenderbuffer($1, $2)", target, renderbuffer)
}

// BindTexture fn
// js:"bindTexture"
func (*WebGLRenderingContext) BindTexture(target uint, texture *webgltexture.WebGLTexture) {
	macro.Rewrite("$_.bindTexture($1, $2)", target, texture)
}

// BlendColor fn
// js:"blendColor"
func (*WebGLRenderingContext) BlendColor(red float32, green float32, blue float32, alpha float32) {
	macro.Rewrite("$_.blendColor($1, $2, $3, $4)", red, green, blue, alpha)
}

// BlendEquation fn
// js:"blendEquation"
func (*WebGLRenderingContext) BlendEquation(mode uint) {
	macro.Rewrite("$_.blendEquation($1)", mode)
}

// BlendEquationSeparate fn
// js:"blendEquationSeparate"
func (*WebGLRenderingContext) BlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	macro.Rewrite("$_.blendEquationSeparate($1, $2)", modeRGB, modeAlpha)
}

// BlendFunc fn
// js:"blendFunc"
func (*WebGLRenderingContext) BlendFunc(sfactor uint, dfactor uint) {
	macro.Rewrite("$_.blendFunc($1, $2)", sfactor, dfactor)
}

// BlendFuncSeparate fn
// js:"blendFuncSeparate"
func (*WebGLRenderingContext) BlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	macro.Rewrite("$_.blendFuncSeparate($1, $2, $3, $4)", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// BufferData fn
// js:"bufferData"
func (*WebGLRenderingContext) BufferData(target uint, size interface{}, usage uint) {
	macro.Rewrite("$_.bufferData($1, $2, $3)", target, size, usage)
}

// BufferSubData fn
// js:"bufferSubData"
func (*WebGLRenderingContext) BufferSubData(target uint, offset int64, data interface{}) {
	macro.Rewrite("$_.bufferSubData($1, $2, $3)", target, offset, data)
}

// CheckFramebufferStatus fn
// js:"checkFramebufferStatus"
func (*WebGLRenderingContext) CheckFramebufferStatus(target uint) (u uint) {
	macro.Rewrite("$_.checkFramebufferStatus($1)", target)
	return u
}

// Clear fn
// js:"clear"
func (*WebGLRenderingContext) Clear(mask uint) {
	macro.Rewrite("$_.clear($1)", mask)
}

// ClearColor fn
// js:"clearColor"
func (*WebGLRenderingContext) ClearColor(red float32, green float32, blue float32, alpha float32) {
	macro.Rewrite("$_.clearColor($1, $2, $3, $4)", red, green, blue, alpha)
}

// ClearDepth fn
// js:"clearDepth"
func (*WebGLRenderingContext) ClearDepth(depth float32) {
	macro.Rewrite("$_.clearDepth($1)", depth)
}

// ClearStencil fn
// js:"clearStencil"
func (*WebGLRenderingContext) ClearStencil(s int) {
	macro.Rewrite("$_.clearStencil($1)", s)
}

// ColorMask fn
// js:"colorMask"
func (*WebGLRenderingContext) ColorMask(red bool, green bool, blue bool, alpha bool) {
	macro.Rewrite("$_.colorMask($1, $2, $3, $4)", red, green, blue, alpha)
}

// CompileShader fn
// js:"compileShader"
func (*WebGLRenderingContext) CompileShader(shader *webglshader.WebGLShader) {
	macro.Rewrite("$_.compileShader($1)", shader)
}

// CompressedTexImage2d fn
// js:"compressedTexImage2D"
func (*WebGLRenderingContext) CompressedTexImage2d(target uint, level int, internalformat uint, width int, height int, border int, data []byte) {
	macro.Rewrite("$_.compressedTexImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, internalformat, width, height, border, data)
}

// CompressedTexSubImage2d fn
// js:"compressedTexSubImage2D"
func (*WebGLRenderingContext) CompressedTexSubImage2d(target uint, level int, xoffset int, yoffset int, width int, height int, format uint, data []byte) {
	macro.Rewrite("$_.compressedTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, width, height, format, data)
}

// CopyTexImage2d fn
// js:"copyTexImage2D"
func (*WebGLRenderingContext) CopyTexImage2d(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	macro.Rewrite("$_.copyTexImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, internalformat, x, y, width, height, border)
}

// CopyTexSubImage2d fn
// js:"copyTexSubImage2D"
func (*WebGLRenderingContext) CopyTexSubImage2d(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	macro.Rewrite("$_.copyTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, x, y, width, height)
}

// CreateBuffer fn
// js:"createBuffer"
func (*WebGLRenderingContext) CreateBuffer() (w *webglbuffer.WebGLBuffer) {
	macro.Rewrite("$_.createBuffer()")
	return w
}

// CreateFramebuffer fn
// js:"createFramebuffer"
func (*WebGLRenderingContext) CreateFramebuffer() (w *webglframebuffer.WebGLFramebuffer) {
	macro.Rewrite("$_.createFramebuffer()")
	return w
}

// CreateProgram fn
// js:"createProgram"
func (*WebGLRenderingContext) CreateProgram() (w *webglprogram.WebGLProgram) {
	macro.Rewrite("$_.createProgram()")
	return w
}

// CreateRenderbuffer fn
// js:"createRenderbuffer"
func (*WebGLRenderingContext) CreateRenderbuffer() (w *webglrenderbuffer.WebGLRenderbuffer) {
	macro.Rewrite("$_.createRenderbuffer()")
	return w
}

// CreateShader fn
// js:"createShader"
func (*WebGLRenderingContext) CreateShader(kind uint) (w *webglshader.WebGLShader) {
	macro.Rewrite("$_.createShader($1)", kind)
	return w
}

// CreateTexture fn
// js:"createTexture"
func (*WebGLRenderingContext) CreateTexture() (w *webgltexture.WebGLTexture) {
	macro.Rewrite("$_.createTexture()")
	return w
}

// CullFace fn
// js:"cullFace"
func (*WebGLRenderingContext) CullFace(mode uint) {
	macro.Rewrite("$_.cullFace($1)", mode)
}

// DeleteBuffer fn
// js:"deleteBuffer"
func (*WebGLRenderingContext) DeleteBuffer(buffer *webglbuffer.WebGLBuffer) {
	macro.Rewrite("$_.deleteBuffer($1)", buffer)
}

// DeleteFramebuffer fn
// js:"deleteFramebuffer"
func (*WebGLRenderingContext) DeleteFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) {
	macro.Rewrite("$_.deleteFramebuffer($1)", framebuffer)
}

// DeleteProgram fn
// js:"deleteProgram"
func (*WebGLRenderingContext) DeleteProgram(program *webglprogram.WebGLProgram) {
	macro.Rewrite("$_.deleteProgram($1)", program)
}

// DeleteRenderbuffer fn
// js:"deleteRenderbuffer"
func (*WebGLRenderingContext) DeleteRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	macro.Rewrite("$_.deleteRenderbuffer($1)", renderbuffer)
}

// DeleteShader fn
// js:"deleteShader"
func (*WebGLRenderingContext) DeleteShader(shader *webglshader.WebGLShader) {
	macro.Rewrite("$_.deleteShader($1)", shader)
}

// DeleteTexture fn
// js:"deleteTexture"
func (*WebGLRenderingContext) DeleteTexture(texture *webgltexture.WebGLTexture) {
	macro.Rewrite("$_.deleteTexture($1)", texture)
}

// DepthFunc fn
// js:"depthFunc"
func (*WebGLRenderingContext) DepthFunc(fn uint) {
	macro.Rewrite("$_.depthFunc($1)", fn)
}

// DepthMask fn
// js:"depthMask"
func (*WebGLRenderingContext) DepthMask(flag bool) {
	macro.Rewrite("$_.depthMask($1)", flag)
}

// DepthRange fn
// js:"depthRange"
func (*WebGLRenderingContext) DepthRange(zNear float32, zFar float32) {
	macro.Rewrite("$_.depthRange($1, $2)", zNear, zFar)
}

// DetachShader fn
// js:"detachShader"
func (*WebGLRenderingContext) DetachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	macro.Rewrite("$_.detachShader($1, $2)", program, shader)
}

// Disable fn
// js:"disable"
func (*WebGLRenderingContext) Disable(capacity uint) {
	macro.Rewrite("$_.disable($1)", capacity)
}

// DisableVertexAttribArray fn
// js:"disableVertexAttribArray"
func (*WebGLRenderingContext) DisableVertexAttribArray(index uint) {
	macro.Rewrite("$_.disableVertexAttribArray($1)", index)
}

// DrawArrays fn
// js:"drawArrays"
func (*WebGLRenderingContext) DrawArrays(mode uint, first int, count int) {
	macro.Rewrite("$_.drawArrays($1, $2, $3)", mode, first, count)
}

// DrawElements fn
// js:"drawElements"
func (*WebGLRenderingContext) DrawElements(mode uint, count int, kind uint, offset int64) {
	macro.Rewrite("$_.drawElements($1, $2, $3, $4)", mode, count, kind, offset)
}

// Enable fn
// js:"enable"
func (*WebGLRenderingContext) Enable(capacity uint) {
	macro.Rewrite("$_.enable($1)", capacity)
}

// EnableVertexAttribArray fn
// js:"enableVertexAttribArray"
func (*WebGLRenderingContext) EnableVertexAttribArray(index uint) {
	macro.Rewrite("$_.enableVertexAttribArray($1)", index)
}

// Finish fn
// js:"finish"
func (*WebGLRenderingContext) Finish() {
	macro.Rewrite("$_.finish()")
}

// Flush fn
// js:"flush"
func (*WebGLRenderingContext) Flush() {
	macro.Rewrite("$_.flush()")
}

// FramebufferRenderbuffer fn
// js:"framebufferRenderbuffer"
func (*WebGLRenderingContext) FramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	macro.Rewrite("$_.framebufferRenderbuffer($1, $2, $3, $4)", target, attachment, renderbuffertarget, renderbuffer)
}

// FramebufferTexture2d fn
// js:"framebufferTexture2D"
func (*WebGLRenderingContext) FramebufferTexture2d(target uint, attachment uint, textarget uint, texture *webgltexture.WebGLTexture, level int) {
	macro.Rewrite("$_.framebufferTexture2D($1, $2, $3, $4, $5)", target, attachment, textarget, texture, level)
}

// FrontFace fn
// js:"frontFace"
func (*WebGLRenderingContext) FrontFace(mode uint) {
	macro.Rewrite("$_.frontFace($1)", mode)
}

// GenerateMipmap fn
// js:"generateMipmap"
func (*WebGLRenderingContext) GenerateMipmap(target uint) {
	macro.Rewrite("$_.generateMipmap($1)", target)
}

// GetActiveAttrib fn
// js:"getActiveAttrib"
func (*WebGLRenderingContext) GetActiveAttrib(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	macro.Rewrite("$_.getActiveAttrib($1, $2)", program, index)
	return w
}

// GetActiveUniform fn
// js:"getActiveUniform"
func (*WebGLRenderingContext) GetActiveUniform(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	macro.Rewrite("$_.getActiveUniform($1, $2)", program, index)
	return w
}

// GetAttachedShaders fn
// js:"getAttachedShaders"
func (*WebGLRenderingContext) GetAttachedShaders(program *webglprogram.WebGLProgram) (w []*webglshader.WebGLShader) {
	macro.Rewrite("$_.getAttachedShaders($1)", program)
	return w
}

// GetAttribLocation fn
// js:"getAttribLocation"
func (*WebGLRenderingContext) GetAttribLocation(program *webglprogram.WebGLProgram, name string) (i int) {
	macro.Rewrite("$_.getAttribLocation($1, $2)", program, name)
	return i
}

// GetBufferParameter fn
// js:"getBufferParameter"
func (*WebGLRenderingContext) GetBufferParameter(target uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getBufferParameter($1, $2)", target, pname)
	return i
}

// GetContextAttributes fn
// js:"getContextAttributes"
func (*WebGLRenderingContext) GetContextAttributes() (w *webglcontextattributes.WebGLContextAttributes) {
	macro.Rewrite("$_.getContextAttributes()")
	return w
}

// GetError fn
// js:"getError"
func (*WebGLRenderingContext) GetError() (u uint) {
	macro.Rewrite("$_.getError()")
	return u
}

// GetExtension fn
// js:"getExtension"
func (*WebGLRenderingContext) GetExtension(name string) (i interface{}) {
	macro.Rewrite("$_.getExtension($1)", name)
	return i
}

// GetFramebufferAttachmentParameter fn
// js:"getFramebufferAttachmentParameter"
func (*WebGLRenderingContext) GetFramebufferAttachmentParameter(target uint, attachment uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getFramebufferAttachmentParameter($1, $2, $3)", target, attachment, pname)
	return i
}

// GetParameter fn
// js:"getParameter"
func (*WebGLRenderingContext) GetParameter(pname uint) (i interface{}) {
	macro.Rewrite("$_.getParameter($1)", pname)
	return i
}

// GetProgramInfoLog fn
// js:"getProgramInfoLog"
func (*WebGLRenderingContext) GetProgramInfoLog(program *webglprogram.WebGLProgram) (s string) {
	macro.Rewrite("$_.getProgramInfoLog($1)", program)
	return s
}

// GetProgramParameter fn
// js:"getProgramParameter"
func (*WebGLRenderingContext) GetProgramParameter(program *webglprogram.WebGLProgram, pname uint) (i interface{}) {
	macro.Rewrite("$_.getProgramParameter($1, $2)", program, pname)
	return i
}

// GetRenderbufferParameter fn
// js:"getRenderbufferParameter"
func (*WebGLRenderingContext) GetRenderbufferParameter(target uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getRenderbufferParameter($1, $2)", target, pname)
	return i
}

// GetShaderInfoLog fn
// js:"getShaderInfoLog"
func (*WebGLRenderingContext) GetShaderInfoLog(shader *webglshader.WebGLShader) (s string) {
	macro.Rewrite("$_.getShaderInfoLog($1)", shader)
	return s
}

// GetShaderParameter fn
// js:"getShaderParameter"
func (*WebGLRenderingContext) GetShaderParameter(shader *webglshader.WebGLShader, pname uint) (i interface{}) {
	macro.Rewrite("$_.getShaderParameter($1, $2)", shader, pname)
	return i
}

// GetShaderPrecisionFormat fn
// js:"getShaderPrecisionFormat"
func (*WebGLRenderingContext) GetShaderPrecisionFormat(shadertype uint, precisiontype uint) (w *webglshaderprecisionformat.WebGLShaderPrecisionFormat) {
	macro.Rewrite("$_.getShaderPrecisionFormat($1, $2)", shadertype, precisiontype)
	return w
}

// GetShaderSource fn
// js:"getShaderSource"
func (*WebGLRenderingContext) GetShaderSource(shader *webglshader.WebGLShader) (s string) {
	macro.Rewrite("$_.getShaderSource($1)", shader)
	return s
}

// GetSupportedExtensions fn
// js:"getSupportedExtensions"
func (*WebGLRenderingContext) GetSupportedExtensions() (s []string) {
	macro.Rewrite("$_.getSupportedExtensions()")
	return s
}

// GetTexParameter fn
// js:"getTexParameter"
func (*WebGLRenderingContext) GetTexParameter(target uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getTexParameter($1, $2)", target, pname)
	return i
}

// GetUniform fn
// js:"getUniform"
func (*WebGLRenderingContext) GetUniform(program *webglprogram.WebGLProgram, location *webgluniformlocation.WebGLUniformLocation) (i interface{}) {
	macro.Rewrite("$_.getUniform($1, $2)", program, location)
	return i
}

// GetUniformLocation fn
// js:"getUniformLocation"
func (*WebGLRenderingContext) GetUniformLocation(program *webglprogram.WebGLProgram, name string) (w *webgluniformlocation.WebGLUniformLocation) {
	macro.Rewrite("$_.getUniformLocation($1, $2)", program, name)
	return w
}

// GetVertexAttrib fn
// js:"getVertexAttrib"
func (*WebGLRenderingContext) GetVertexAttrib(index uint, pname uint) (i interface{}) {
	macro.Rewrite("$_.getVertexAttrib($1, $2)", index, pname)
	return i
}

// GetVertexAttribOffset fn
// js:"getVertexAttribOffset"
func (*WebGLRenderingContext) GetVertexAttribOffset(index uint, pname uint) (i int64) {
	macro.Rewrite("$_.getVertexAttribOffset($1, $2)", index, pname)
	return i
}

// Hint fn
// js:"hint"
func (*WebGLRenderingContext) Hint(target uint, mode uint) {
	macro.Rewrite("$_.hint($1, $2)", target, mode)
}

// IsBuffer fn
// js:"isBuffer"
func (*WebGLRenderingContext) IsBuffer(buffer *webglbuffer.WebGLBuffer) (b bool) {
	macro.Rewrite("$_.isBuffer($1)", buffer)
	return b
}

// IsContextLost fn
// js:"isContextLost"
func (*WebGLRenderingContext) IsContextLost() (b bool) {
	macro.Rewrite("$_.isContextLost()")
	return b
}

// IsEnabled fn
// js:"isEnabled"
func (*WebGLRenderingContext) IsEnabled(capacity uint) (b bool) {
	macro.Rewrite("$_.isEnabled($1)", capacity)
	return b
}

// IsFramebuffer fn
// js:"isFramebuffer"
func (*WebGLRenderingContext) IsFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) (b bool) {
	macro.Rewrite("$_.isFramebuffer($1)", framebuffer)
	return b
}

// IsProgram fn
// js:"isProgram"
func (*WebGLRenderingContext) IsProgram(program *webglprogram.WebGLProgram) (b bool) {
	macro.Rewrite("$_.isProgram($1)", program)
	return b
}

// IsRenderbuffer fn
// js:"isRenderbuffer"
func (*WebGLRenderingContext) IsRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) (b bool) {
	macro.Rewrite("$_.isRenderbuffer($1)", renderbuffer)
	return b
}

// IsShader fn
// js:"isShader"
func (*WebGLRenderingContext) IsShader(shader *webglshader.WebGLShader) (b bool) {
	macro.Rewrite("$_.isShader($1)", shader)
	return b
}

// IsTexture fn
// js:"isTexture"
func (*WebGLRenderingContext) IsTexture(texture *webgltexture.WebGLTexture) (b bool) {
	macro.Rewrite("$_.isTexture($1)", texture)
	return b
}

// LineWidth fn
// js:"lineWidth"
func (*WebGLRenderingContext) LineWidth(width float32) {
	macro.Rewrite("$_.lineWidth($1)", width)
}

// LinkProgram fn
// js:"linkProgram"
func (*WebGLRenderingContext) LinkProgram(program *webglprogram.WebGLProgram) {
	macro.Rewrite("$_.linkProgram($1)", program)
}

// PixelStorei fn
// js:"pixelStorei"
func (*WebGLRenderingContext) PixelStorei(pname uint, param int) {
	macro.Rewrite("$_.pixelStorei($1, $2)", pname, param)
}

// PolygonOffset fn
// js:"polygonOffset"
func (*WebGLRenderingContext) PolygonOffset(factor float32, units float32) {
	macro.Rewrite("$_.polygonOffset($1, $2)", factor, units)
}

// ReadPixels fn
// js:"readPixels"
func (*WebGLRenderingContext) ReadPixels(x int, y int, width int, height int, format uint, kind uint, pixels []byte) {
	macro.Rewrite("$_.readPixels($1, $2, $3, $4, $5, $6, $7)", x, y, width, height, format, kind, pixels)
}

// RenderbufferStorage fn
// js:"renderbufferStorage"
func (*WebGLRenderingContext) RenderbufferStorage(target uint, internalformat uint, width int, height int) {
	macro.Rewrite("$_.renderbufferStorage($1, $2, $3, $4)", target, internalformat, width, height)
}

// SampleCoverage fn
// js:"sampleCoverage"
func (*WebGLRenderingContext) SampleCoverage(value float32, invert bool) {
	macro.Rewrite("$_.sampleCoverage($1, $2)", value, invert)
}

// Scissor fn
// js:"scissor"
func (*WebGLRenderingContext) Scissor(x int, y int, width int, height int) {
	macro.Rewrite("$_.scissor($1, $2, $3, $4)", x, y, width, height)
}

// ShaderSource fn
// js:"shaderSource"
func (*WebGLRenderingContext) ShaderSource(shader *webglshader.WebGLShader, source string) {
	macro.Rewrite("$_.shaderSource($1, $2)", shader, source)
}

// StencilFunc fn
// js:"stencilFunc"
func (*WebGLRenderingContext) StencilFunc(fn uint, ref int, mask uint) {
	macro.Rewrite("$_.stencilFunc($1, $2, $3)", fn, ref, mask)
}

// StencilFuncSeparate fn
// js:"stencilFuncSeparate"
func (*WebGLRenderingContext) StencilFuncSeparate(face uint, fn uint, ref int, mask uint) {
	macro.Rewrite("$_.stencilFuncSeparate($1, $2, $3, $4)", face, fn, ref, mask)
}

// StencilMask fn
// js:"stencilMask"
func (*WebGLRenderingContext) StencilMask(mask uint) {
	macro.Rewrite("$_.stencilMask($1)", mask)
}

// StencilMaskSeparate fn
// js:"stencilMaskSeparate"
func (*WebGLRenderingContext) StencilMaskSeparate(face uint, mask uint) {
	macro.Rewrite("$_.stencilMaskSeparate($1, $2)", face, mask)
}

// StencilOp fn
// js:"stencilOp"
func (*WebGLRenderingContext) StencilOp(fail uint, zfail uint, zpass uint) {
	macro.Rewrite("$_.stencilOp($1, $2, $3)", fail, zfail, zpass)
}

// StencilOpSeparate fn
// js:"stencilOpSeparate"
func (*WebGLRenderingContext) StencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	macro.Rewrite("$_.stencilOpSeparate($1, $2, $3, $4)", face, fail, zfail, zpass)
}

// TexImage2d fn
// js:"texImage2D"
func (*WebGLRenderingContext) TexImage2d(target uint, level int, internalformat uint, format uint, kind uint, pixels *imagedata.ImageData) {
	macro.Rewrite("$_.texImage2D($1, $2, $3, $4, $5, $6)", target, level, internalformat, format, kind, pixels)
}

// TexParameterf fn
// js:"texParameterf"
func (*WebGLRenderingContext) TexParameterf(target uint, pname uint, param float32) {
	macro.Rewrite("$_.texParameterf($1, $2, $3)", target, pname, param)
}

// TexParameteri fn
// js:"texParameteri"
func (*WebGLRenderingContext) TexParameteri(target uint, pname uint, param int) {
	macro.Rewrite("$_.texParameteri($1, $2, $3)", target, pname, param)
}

// TexSubImage2d fn
// js:"texSubImage2D"
func (*WebGLRenderingContext) TexSubImage2d(target uint, level int, xoffset int, yoffset int, format uint, kind uint, pixels *imagedata.ImageData) {
	macro.Rewrite("$_.texSubImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, xoffset, yoffset, format, kind, pixels)
}

// Uniform1f fn
// js:"uniform1f"
func (*WebGLRenderingContext) Uniform1f(location *webgluniformlocation.WebGLUniformLocation, x float32) {
	macro.Rewrite("$_.uniform1f($1, $2)", location, x)
}

// Uniform1fv fn
// js:"uniform1fv"
func (*WebGLRenderingContext) Uniform1fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform1fv($1, $2)", location, v)
}

// Uniform1i fn
// js:"uniform1i"
func (*WebGLRenderingContext) Uniform1i(location *webgluniformlocation.WebGLUniformLocation, x int) {
	macro.Rewrite("$_.uniform1i($1, $2)", location, x)
}

// Uniform1iv fn
// js:"uniform1iv"
func (*WebGLRenderingContext) Uniform1iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform1iv($1, $2)", location, v)
}

// Uniform2f fn
// js:"uniform2f"
func (*WebGLRenderingContext) Uniform2f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32) {
	macro.Rewrite("$_.uniform2f($1, $2, $3)", location, x, y)
}

// Uniform2fv fn
// js:"uniform2fv"
func (*WebGLRenderingContext) Uniform2fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform2fv($1, $2)", location, v)
}

// Uniform2i fn
// js:"uniform2i"
func (*WebGLRenderingContext) Uniform2i(location *webgluniformlocation.WebGLUniformLocation, x int, y int) {
	macro.Rewrite("$_.uniform2i($1, $2, $3)", location, x, y)
}

// Uniform2iv fn
// js:"uniform2iv"
func (*WebGLRenderingContext) Uniform2iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform2iv($1, $2)", location, v)
}

// Uniform3f fn
// js:"uniform3f"
func (*WebGLRenderingContext) Uniform3f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32) {
	macro.Rewrite("$_.uniform3f($1, $2, $3, $4)", location, x, y, z)
}

// Uniform3fv fn
// js:"uniform3fv"
func (*WebGLRenderingContext) Uniform3fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform3fv($1, $2)", location, v)
}

// Uniform3i fn
// js:"uniform3i"
func (*WebGLRenderingContext) Uniform3i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int) {
	macro.Rewrite("$_.uniform3i($1, $2, $3, $4)", location, x, y, z)
}

// Uniform3iv fn
// js:"uniform3iv"
func (*WebGLRenderingContext) Uniform3iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform3iv($1, $2)", location, v)
}

// Uniform4f fn
// js:"uniform4f"
func (*WebGLRenderingContext) Uniform4f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	macro.Rewrite("$_.uniform4f($1, $2, $3, $4, $5)", location, x, y, z, w)
}

// Uniform4fv fn
// js:"uniform4fv"
func (*WebGLRenderingContext) Uniform4fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	macro.Rewrite("$_.uniform4fv($1, $2)", location, v)
}

// Uniform4i fn
// js:"uniform4i"
func (*WebGLRenderingContext) Uniform4i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int, w int) {
	macro.Rewrite("$_.uniform4i($1, $2, $3, $4, $5)", location, x, y, z, w)
}

// Uniform4iv fn
// js:"uniform4iv"
func (*WebGLRenderingContext) Uniform4iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	macro.Rewrite("$_.uniform4iv($1, $2)", location, v)
}

// UniformMatrix2fv fn
// js:"uniformMatrix2fv"
func (*WebGLRenderingContext) UniformMatrix2fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	macro.Rewrite("$_.uniformMatrix2fv($1, $2, $3)", location, transpose, value)
}

// UniformMatrix3fv fn
// js:"uniformMatrix3fv"
func (*WebGLRenderingContext) UniformMatrix3fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	macro.Rewrite("$_.uniformMatrix3fv($1, $2, $3)", location, transpose, value)
}

// UniformMatrix4fv fn
// js:"uniformMatrix4fv"
func (*WebGLRenderingContext) UniformMatrix4fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	macro.Rewrite("$_.uniformMatrix4fv($1, $2, $3)", location, transpose, value)
}

// UseProgram fn
// js:"useProgram"
func (*WebGLRenderingContext) UseProgram(program *webglprogram.WebGLProgram) {
	macro.Rewrite("$_.useProgram($1)", program)
}

// ValidateProgram fn
// js:"validateProgram"
func (*WebGLRenderingContext) ValidateProgram(program *webglprogram.WebGLProgram) {
	macro.Rewrite("$_.validateProgram($1)", program)
}

// VertexAttrib1f fn
// js:"vertexAttrib1f"
func (*WebGLRenderingContext) VertexAttrib1f(indx uint, x float32) {
	macro.Rewrite("$_.vertexAttrib1f($1, $2)", indx, x)
}

// VertexAttrib1fv fn
// js:"vertexAttrib1fv"
func (*WebGLRenderingContext) VertexAttrib1fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib1fv($1, $2)", indx, values)
}

// VertexAttrib2f fn
// js:"vertexAttrib2f"
func (*WebGLRenderingContext) VertexAttrib2f(indx uint, x float32, y float32) {
	macro.Rewrite("$_.vertexAttrib2f($1, $2, $3)", indx, x, y)
}

// VertexAttrib2fv fn
// js:"vertexAttrib2fv"
func (*WebGLRenderingContext) VertexAttrib2fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib2fv($1, $2)", indx, values)
}

// VertexAttrib3f fn
// js:"vertexAttrib3f"
func (*WebGLRenderingContext) VertexAttrib3f(indx uint, x float32, y float32, z float32) {
	macro.Rewrite("$_.vertexAttrib3f($1, $2, $3, $4)", indx, x, y, z)
}

// VertexAttrib3fv fn
// js:"vertexAttrib3fv"
func (*WebGLRenderingContext) VertexAttrib3fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib3fv($1, $2)", indx, values)
}

// VertexAttrib4f fn
// js:"vertexAttrib4f"
func (*WebGLRenderingContext) VertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	macro.Rewrite("$_.vertexAttrib4f($1, $2, $3, $4, $5)", indx, x, y, z, w)
}

// VertexAttrib4fv fn
// js:"vertexAttrib4fv"
func (*WebGLRenderingContext) VertexAttrib4fv(indx uint, values []float32) {
	macro.Rewrite("$_.vertexAttrib4fv($1, $2)", indx, values)
}

// VertexAttribPointer fn
// js:"vertexAttribPointer"
func (*WebGLRenderingContext) VertexAttribPointer(indx uint, size int, kind uint, normalized bool, stride int, offset int64) {
	macro.Rewrite("$_.vertexAttribPointer($1, $2, $3, $4, $5, $6)", indx, size, kind, normalized, stride, offset)
}

// Viewport fn
// js:"viewport"
func (*WebGLRenderingContext) Viewport(x int, y int, width int, height int) {
	macro.Rewrite("$_.viewport($1, $2, $3, $4)", x, y, width, height)
}

// Canvas prop
// js:"canvas"
func (*WebGLRenderingContext) Canvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	macro.Rewrite("$_.canvas")
	return canvas
}

// DrawingBufferHeight prop
// js:"drawingBufferHeight"
func (*WebGLRenderingContext) DrawingBufferHeight() (drawingBufferHeight int) {
	macro.Rewrite("$_.drawingBufferHeight")
	return drawingBufferHeight
}

// DrawingBufferWidth prop
// js:"drawingBufferWidth"
func (*WebGLRenderingContext) DrawingBufferWidth() (drawingBufferWidth int) {
	macro.Rewrite("$_.drawingBufferWidth")
	return drawingBufferWidth
}
