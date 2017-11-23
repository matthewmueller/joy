package webglrenderingcontext

import (
	"github.com/matthewmueller/golly/dom2/htmlcanvaselement"
	"github.com/matthewmueller/golly/dom2/imagedata"
	"github.com/matthewmueller/golly/dom2/webglactiveinfo"
	"github.com/matthewmueller/golly/dom2/webglbuffer"
	"github.com/matthewmueller/golly/dom2/webglcontextattributes"
	"github.com/matthewmueller/golly/dom2/webglframebuffer"
	"github.com/matthewmueller/golly/dom2/webglprogram"
	"github.com/matthewmueller/golly/dom2/webglrenderbuffer"
	"github.com/matthewmueller/golly/dom2/webglshader"
	"github.com/matthewmueller/golly/dom2/webglshaderprecisionformat"
	"github.com/matthewmueller/golly/dom2/webgltexture"
	"github.com/matthewmueller/golly/dom2/webgluniformlocation"
	"github.com/matthewmueller/golly/js"
)

// WebGLRenderingContext struct
// js:"WebGLRenderingContext,omit"
type WebGLRenderingContext struct {
}

// ActiveTexture fn
func (*WebGLRenderingContext) ActiveTexture(texture uint) {
	js.Rewrite("$<.activeTexture($1)", texture)
}

// AttachShader fn
func (*WebGLRenderingContext) AttachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	js.Rewrite("$<.attachShader($1, $2)", program, shader)
}

// BindAttribLocation fn
func (*WebGLRenderingContext) BindAttribLocation(program *webglprogram.WebGLProgram, index uint, name string) {
	js.Rewrite("$<.bindAttribLocation($1, $2, $3)", program, index, name)
}

// BindBuffer fn
func (*WebGLRenderingContext) BindBuffer(target uint, buffer *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.bindBuffer($1, $2)", target, buffer)
}

// BindFramebuffer fn
func (*WebGLRenderingContext) BindFramebuffer(target uint, framebuffer *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.bindFramebuffer($1, $2)", target, framebuffer)
}

// BindRenderbuffer fn
func (*WebGLRenderingContext) BindRenderbuffer(target uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.bindRenderbuffer($1, $2)", target, renderbuffer)
}

// BindTexture fn
func (*WebGLRenderingContext) BindTexture(target uint, texture *webgltexture.WebGLTexture) {
	js.Rewrite("$<.bindTexture($1, $2)", target, texture)
}

// BlendColor fn
func (*WebGLRenderingContext) BlendColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.blendColor($1, $2, $3, $4)", red, green, blue, alpha)
}

// BlendEquation fn
func (*WebGLRenderingContext) BlendEquation(mode uint) {
	js.Rewrite("$<.blendEquation($1)", mode)
}

// BlendEquationSeparate fn
func (*WebGLRenderingContext) BlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	js.Rewrite("$<.blendEquationSeparate($1, $2)", modeRGB, modeAlpha)
}

// BlendFunc fn
func (*WebGLRenderingContext) BlendFunc(sfactor uint, dfactor uint) {
	js.Rewrite("$<.blendFunc($1, $2)", sfactor, dfactor)
}

// BlendFuncSeparate fn
func (*WebGLRenderingContext) BlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	js.Rewrite("$<.blendFuncSeparate($1, $2, $3, $4)", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// BufferData fn
func (*WebGLRenderingContext) BufferData(target uint, size interface{}, usage uint) {
	js.Rewrite("$<.bufferData($1, $2, $3)", target, size, usage)
}

// BufferSubData fn
func (*WebGLRenderingContext) BufferSubData(target uint, offset int64, data interface{}) {
	js.Rewrite("$<.bufferSubData($1, $2, $3)", target, offset, data)
}

// CheckFramebufferStatus fn
func (*WebGLRenderingContext) CheckFramebufferStatus(target uint) (u uint) {
	js.Rewrite("$<.checkFramebufferStatus($1)", target)
	return u
}

// Clear fn
func (*WebGLRenderingContext) Clear(mask uint) {
	js.Rewrite("$<.clear($1)", mask)
}

// ClearColor fn
func (*WebGLRenderingContext) ClearColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.clearColor($1, $2, $3, $4)", red, green, blue, alpha)
}

// ClearDepth fn
func (*WebGLRenderingContext) ClearDepth(depth float32) {
	js.Rewrite("$<.clearDepth($1)", depth)
}

// ClearStencil fn
func (*WebGLRenderingContext) ClearStencil(s int) {
	js.Rewrite("$<.clearStencil($1)", s)
}

// ColorMask fn
func (*WebGLRenderingContext) ColorMask(red bool, green bool, blue bool, alpha bool) {
	js.Rewrite("$<.colorMask($1, $2, $3, $4)", red, green, blue, alpha)
}

// CompileShader fn
func (*WebGLRenderingContext) CompileShader(shader *webglshader.WebGLShader) {
	js.Rewrite("$<.compileShader($1)", shader)
}

// CompressedTexImage2d fn
func (*WebGLRenderingContext) CompressedTexImage2d(target uint, level int, internalformat uint, width int, height int, border int, data []byte) {
	js.Rewrite("$<.compressedTexImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, internalformat, width, height, border, data)
}

// CompressedTexSubImage2d fn
func (*WebGLRenderingContext) CompressedTexSubImage2d(target uint, level int, xoffset int, yoffset int, width int, height int, format uint, data []byte) {
	js.Rewrite("$<.compressedTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, width, height, format, data)
}

// CopyTexImage2d fn
func (*WebGLRenderingContext) CopyTexImage2d(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	js.Rewrite("$<.copyTexImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, internalformat, x, y, width, height, border)
}

// CopyTexSubImage2d fn
func (*WebGLRenderingContext) CopyTexSubImage2d(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	js.Rewrite("$<.copyTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, x, y, width, height)
}

// CreateBuffer fn
func (*WebGLRenderingContext) CreateBuffer() (w *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.createBuffer()")
	return w
}

// CreateFramebuffer fn
func (*WebGLRenderingContext) CreateFramebuffer() (w *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.createFramebuffer()")
	return w
}

// CreateProgram fn
func (*WebGLRenderingContext) CreateProgram() (w *webglprogram.WebGLProgram) {
	js.Rewrite("$<.createProgram()")
	return w
}

// CreateRenderbuffer fn
func (*WebGLRenderingContext) CreateRenderbuffer() (w *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.createRenderbuffer()")
	return w
}

// CreateShader fn
func (*WebGLRenderingContext) CreateShader(kind uint) (w *webglshader.WebGLShader) {
	js.Rewrite("$<.createShader($1)", kind)
	return w
}

// CreateTexture fn
func (*WebGLRenderingContext) CreateTexture() (w *webgltexture.WebGLTexture) {
	js.Rewrite("$<.createTexture()")
	return w
}

// CullFace fn
func (*WebGLRenderingContext) CullFace(mode uint) {
	js.Rewrite("$<.cullFace($1)", mode)
}

// DeleteBuffer fn
func (*WebGLRenderingContext) DeleteBuffer(buffer *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.deleteBuffer($1)", buffer)
}

// DeleteFramebuffer fn
func (*WebGLRenderingContext) DeleteFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.deleteFramebuffer($1)", framebuffer)
}

// DeleteProgram fn
func (*WebGLRenderingContext) DeleteProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.deleteProgram($1)", program)
}

// DeleteRenderbuffer fn
func (*WebGLRenderingContext) DeleteRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.deleteRenderbuffer($1)", renderbuffer)
}

// DeleteShader fn
func (*WebGLRenderingContext) DeleteShader(shader *webglshader.WebGLShader) {
	js.Rewrite("$<.deleteShader($1)", shader)
}

// DeleteTexture fn
func (*WebGLRenderingContext) DeleteTexture(texture *webgltexture.WebGLTexture) {
	js.Rewrite("$<.deleteTexture($1)", texture)
}

// DepthFunc fn
func (*WebGLRenderingContext) DepthFunc(fn uint) {
	js.Rewrite("$<.depthFunc($1)", fn)
}

// DepthMask fn
func (*WebGLRenderingContext) DepthMask(flag bool) {
	js.Rewrite("$<.depthMask($1)", flag)
}

// DepthRange fn
func (*WebGLRenderingContext) DepthRange(zNear float32, zFar float32) {
	js.Rewrite("$<.depthRange($1, $2)", zNear, zFar)
}

// DetachShader fn
func (*WebGLRenderingContext) DetachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	js.Rewrite("$<.detachShader($1, $2)", program, shader)
}

// Disable fn
func (*WebGLRenderingContext) Disable(capacity uint) {
	js.Rewrite("$<.disable($1)", capacity)
}

// DisableVertexAttribArray fn
func (*WebGLRenderingContext) DisableVertexAttribArray(index uint) {
	js.Rewrite("$<.disableVertexAttribArray($1)", index)
}

// DrawArrays fn
func (*WebGLRenderingContext) DrawArrays(mode uint, first int, count int) {
	js.Rewrite("$<.drawArrays($1, $2, $3)", mode, first, count)
}

// DrawElements fn
func (*WebGLRenderingContext) DrawElements(mode uint, count int, kind uint, offset int64) {
	js.Rewrite("$<.drawElements($1, $2, $3, $4)", mode, count, kind, offset)
}

// Enable fn
func (*WebGLRenderingContext) Enable(capacity uint) {
	js.Rewrite("$<.enable($1)", capacity)
}

// EnableVertexAttribArray fn
func (*WebGLRenderingContext) EnableVertexAttribArray(index uint) {
	js.Rewrite("$<.enableVertexAttribArray($1)", index)
}

// Finish fn
func (*WebGLRenderingContext) Finish() {
	js.Rewrite("$<.finish()")
}

// Flush fn
func (*WebGLRenderingContext) Flush() {
	js.Rewrite("$<.flush()")
}

// FramebufferRenderbuffer fn
func (*WebGLRenderingContext) FramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.framebufferRenderbuffer($1, $2, $3, $4)", target, attachment, renderbuffertarget, renderbuffer)
}

// FramebufferTexture2d fn
func (*WebGLRenderingContext) FramebufferTexture2d(target uint, attachment uint, textarget uint, texture *webgltexture.WebGLTexture, level int) {
	js.Rewrite("$<.framebufferTexture2D($1, $2, $3, $4, $5)", target, attachment, textarget, texture, level)
}

// FrontFace fn
func (*WebGLRenderingContext) FrontFace(mode uint) {
	js.Rewrite("$<.frontFace($1)", mode)
}

// GenerateMipmap fn
func (*WebGLRenderingContext) GenerateMipmap(target uint) {
	js.Rewrite("$<.generateMipmap($1)", target)
}

// GetActiveAttrib fn
func (*WebGLRenderingContext) GetActiveAttrib(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	js.Rewrite("$<.getActiveAttrib($1, $2)", program, index)
	return w
}

// GetActiveUniform fn
func (*WebGLRenderingContext) GetActiveUniform(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	js.Rewrite("$<.getActiveUniform($1, $2)", program, index)
	return w
}

// GetAttachedShaders fn
func (*WebGLRenderingContext) GetAttachedShaders(program *webglprogram.WebGLProgram) (w []*webglshader.WebGLShader) {
	js.Rewrite("$<.getAttachedShaders($1)", program)
	return w
}

// GetAttribLocation fn
func (*WebGLRenderingContext) GetAttribLocation(program *webglprogram.WebGLProgram, name string) (i int) {
	js.Rewrite("$<.getAttribLocation($1, $2)", program, name)
	return i
}

// GetBufferParameter fn
func (*WebGLRenderingContext) GetBufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getBufferParameter($1, $2)", target, pname)
	return i
}

// GetContextAttributes fn
func (*WebGLRenderingContext) GetContextAttributes() (w *webglcontextattributes.WebGLContextAttributes) {
	js.Rewrite("$<.getContextAttributes()")
	return w
}

// GetError fn
func (*WebGLRenderingContext) GetError() (u uint) {
	js.Rewrite("$<.getError()")
	return u
}

// GetExtension fn
func (*WebGLRenderingContext) GetExtension(name string) (i interface{}) {
	js.Rewrite("$<.getExtension($1)", name)
	return i
}

// GetFramebufferAttachmentParameter fn
func (*WebGLRenderingContext) GetFramebufferAttachmentParameter(target uint, attachment uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getFramebufferAttachmentParameter($1, $2, $3)", target, attachment, pname)
	return i
}

// GetParameter fn
func (*WebGLRenderingContext) GetParameter(pname uint) (i interface{}) {
	js.Rewrite("$<.getParameter($1)", pname)
	return i
}

// GetProgramInfoLog fn
func (*WebGLRenderingContext) GetProgramInfoLog(program *webglprogram.WebGLProgram) (s string) {
	js.Rewrite("$<.getProgramInfoLog($1)", program)
	return s
}

// GetProgramParameter fn
func (*WebGLRenderingContext) GetProgramParameter(program *webglprogram.WebGLProgram, pname uint) (i interface{}) {
	js.Rewrite("$<.getProgramParameter($1, $2)", program, pname)
	return i
}

// GetRenderbufferParameter fn
func (*WebGLRenderingContext) GetRenderbufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getRenderbufferParameter($1, $2)", target, pname)
	return i
}

// GetShaderInfoLog fn
func (*WebGLRenderingContext) GetShaderInfoLog(shader *webglshader.WebGLShader) (s string) {
	js.Rewrite("$<.getShaderInfoLog($1)", shader)
	return s
}

// GetShaderParameter fn
func (*WebGLRenderingContext) GetShaderParameter(shader *webglshader.WebGLShader, pname uint) (i interface{}) {
	js.Rewrite("$<.getShaderParameter($1, $2)", shader, pname)
	return i
}

// GetShaderPrecisionFormat fn
func (*WebGLRenderingContext) GetShaderPrecisionFormat(shadertype uint, precisiontype uint) (w *webglshaderprecisionformat.WebGLShaderPrecisionFormat) {
	js.Rewrite("$<.getShaderPrecisionFormat($1, $2)", shadertype, precisiontype)
	return w
}

// GetShaderSource fn
func (*WebGLRenderingContext) GetShaderSource(shader *webglshader.WebGLShader) (s string) {
	js.Rewrite("$<.getShaderSource($1)", shader)
	return s
}

// GetSupportedExtensions fn
func (*WebGLRenderingContext) GetSupportedExtensions() (s []string) {
	js.Rewrite("$<.getSupportedExtensions()")
	return s
}

// GetTexParameter fn
func (*WebGLRenderingContext) GetTexParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getTexParameter($1, $2)", target, pname)
	return i
}

// GetUniform fn
func (*WebGLRenderingContext) GetUniform(program *webglprogram.WebGLProgram, location *webgluniformlocation.WebGLUniformLocation) (i interface{}) {
	js.Rewrite("$<.getUniform($1, $2)", program, location)
	return i
}

// GetUniformLocation fn
func (*WebGLRenderingContext) GetUniformLocation(program *webglprogram.WebGLProgram, name string) (w *webgluniformlocation.WebGLUniformLocation) {
	js.Rewrite("$<.getUniformLocation($1, $2)", program, name)
	return w
}

// GetVertexAttrib fn
func (*WebGLRenderingContext) GetVertexAttrib(index uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getVertexAttrib($1, $2)", index, pname)
	return i
}

// GetVertexAttribOffset fn
func (*WebGLRenderingContext) GetVertexAttribOffset(index uint, pname uint) (i int64) {
	js.Rewrite("$<.getVertexAttribOffset($1, $2)", index, pname)
	return i
}

// Hint fn
func (*WebGLRenderingContext) Hint(target uint, mode uint) {
	js.Rewrite("$<.hint($1, $2)", target, mode)
}

// IsBuffer fn
func (*WebGLRenderingContext) IsBuffer(buffer *webglbuffer.WebGLBuffer) (b bool) {
	js.Rewrite("$<.isBuffer($1)", buffer)
	return b
}

// IsContextLost fn
func (*WebGLRenderingContext) IsContextLost() (b bool) {
	js.Rewrite("$<.isContextLost()")
	return b
}

// IsEnabled fn
func (*WebGLRenderingContext) IsEnabled(capacity uint) (b bool) {
	js.Rewrite("$<.isEnabled($1)", capacity)
	return b
}

// IsFramebuffer fn
func (*WebGLRenderingContext) IsFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) (b bool) {
	js.Rewrite("$<.isFramebuffer($1)", framebuffer)
	return b
}

// IsProgram fn
func (*WebGLRenderingContext) IsProgram(program *webglprogram.WebGLProgram) (b bool) {
	js.Rewrite("$<.isProgram($1)", program)
	return b
}

// IsRenderbuffer fn
func (*WebGLRenderingContext) IsRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) (b bool) {
	js.Rewrite("$<.isRenderbuffer($1)", renderbuffer)
	return b
}

// IsShader fn
func (*WebGLRenderingContext) IsShader(shader *webglshader.WebGLShader) (b bool) {
	js.Rewrite("$<.isShader($1)", shader)
	return b
}

// IsTexture fn
func (*WebGLRenderingContext) IsTexture(texture *webgltexture.WebGLTexture) (b bool) {
	js.Rewrite("$<.isTexture($1)", texture)
	return b
}

// LineWidth fn
func (*WebGLRenderingContext) LineWidth(width float32) {
	js.Rewrite("$<.lineWidth($1)", width)
}

// LinkProgram fn
func (*WebGLRenderingContext) LinkProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.linkProgram($1)", program)
}

// PixelStorei fn
func (*WebGLRenderingContext) PixelStorei(pname uint, param int) {
	js.Rewrite("$<.pixelStorei($1, $2)", pname, param)
}

// PolygonOffset fn
func (*WebGLRenderingContext) PolygonOffset(factor float32, units float32) {
	js.Rewrite("$<.polygonOffset($1, $2)", factor, units)
}

// ReadPixels fn
func (*WebGLRenderingContext) ReadPixels(x int, y int, width int, height int, format uint, kind uint, pixels []byte) {
	js.Rewrite("$<.readPixels($1, $2, $3, $4, $5, $6, $7)", x, y, width, height, format, kind, pixels)
}

// RenderbufferStorage fn
func (*WebGLRenderingContext) RenderbufferStorage(target uint, internalformat uint, width int, height int) {
	js.Rewrite("$<.renderbufferStorage($1, $2, $3, $4)", target, internalformat, width, height)
}

// SampleCoverage fn
func (*WebGLRenderingContext) SampleCoverage(value float32, invert bool) {
	js.Rewrite("$<.sampleCoverage($1, $2)", value, invert)
}

// Scissor fn
func (*WebGLRenderingContext) Scissor(x int, y int, width int, height int) {
	js.Rewrite("$<.scissor($1, $2, $3, $4)", x, y, width, height)
}

// ShaderSource fn
func (*WebGLRenderingContext) ShaderSource(shader *webglshader.WebGLShader, source string) {
	js.Rewrite("$<.shaderSource($1, $2)", shader, source)
}

// StencilFunc fn
func (*WebGLRenderingContext) StencilFunc(fn uint, ref int, mask uint) {
	js.Rewrite("$<.stencilFunc($1, $2, $3)", fn, ref, mask)
}

// StencilFuncSeparate fn
func (*WebGLRenderingContext) StencilFuncSeparate(face uint, fn uint, ref int, mask uint) {
	js.Rewrite("$<.stencilFuncSeparate($1, $2, $3, $4)", face, fn, ref, mask)
}

// StencilMask fn
func (*WebGLRenderingContext) StencilMask(mask uint) {
	js.Rewrite("$<.stencilMask($1)", mask)
}

// StencilMaskSeparate fn
func (*WebGLRenderingContext) StencilMaskSeparate(face uint, mask uint) {
	js.Rewrite("$<.stencilMaskSeparate($1, $2)", face, mask)
}

// StencilOp fn
func (*WebGLRenderingContext) StencilOp(fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.stencilOp($1, $2, $3)", fail, zfail, zpass)
}

// StencilOpSeparate fn
func (*WebGLRenderingContext) StencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.stencilOpSeparate($1, $2, $3, $4)", face, fail, zfail, zpass)
}

// TexImage2d fn
func (*WebGLRenderingContext) TexImage2d(target uint, level int, internalformat uint, format uint, kind uint, pixels *imagedata.ImageData) {
	js.Rewrite("$<.texImage2D($1, $2, $3, $4, $5, $6)", target, level, internalformat, format, kind, pixels)
}

// TexParameterf fn
func (*WebGLRenderingContext) TexParameterf(target uint, pname uint, param float32) {
	js.Rewrite("$<.texParameterf($1, $2, $3)", target, pname, param)
}

// TexParameteri fn
func (*WebGLRenderingContext) TexParameteri(target uint, pname uint, param int) {
	js.Rewrite("$<.texParameteri($1, $2, $3)", target, pname, param)
}

// TexSubImage2d fn
func (*WebGLRenderingContext) TexSubImage2d(target uint, level int, xoffset int, yoffset int, format uint, kind uint, pixels *imagedata.ImageData) {
	js.Rewrite("$<.texSubImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, xoffset, yoffset, format, kind, pixels)
}

// Uniform1f fn
func (*WebGLRenderingContext) Uniform1f(location *webgluniformlocation.WebGLUniformLocation, x float32) {
	js.Rewrite("$<.uniform1f($1, $2)", location, x)
}

// Uniform1fv fn
func (*WebGLRenderingContext) Uniform1fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform1fv($1, $2)", location, v)
}

// Uniform1i fn
func (*WebGLRenderingContext) Uniform1i(location *webgluniformlocation.WebGLUniformLocation, x int) {
	js.Rewrite("$<.uniform1i($1, $2)", location, x)
}

// Uniform1iv fn
func (*WebGLRenderingContext) Uniform1iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform1iv($1, $2)", location, v)
}

// Uniform2f fn
func (*WebGLRenderingContext) Uniform2f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32) {
	js.Rewrite("$<.uniform2f($1, $2, $3)", location, x, y)
}

// Uniform2fv fn
func (*WebGLRenderingContext) Uniform2fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform2fv($1, $2)", location, v)
}

// Uniform2i fn
func (*WebGLRenderingContext) Uniform2i(location *webgluniformlocation.WebGLUniformLocation, x int, y int) {
	js.Rewrite("$<.uniform2i($1, $2, $3)", location, x, y)
}

// Uniform2iv fn
func (*WebGLRenderingContext) Uniform2iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform2iv($1, $2)", location, v)
}

// Uniform3f fn
func (*WebGLRenderingContext) Uniform3f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32) {
	js.Rewrite("$<.uniform3f($1, $2, $3, $4)", location, x, y, z)
}

// Uniform3fv fn
func (*WebGLRenderingContext) Uniform3fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform3fv($1, $2)", location, v)
}

// Uniform3i fn
func (*WebGLRenderingContext) Uniform3i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int) {
	js.Rewrite("$<.uniform3i($1, $2, $3, $4)", location, x, y, z)
}

// Uniform3iv fn
func (*WebGLRenderingContext) Uniform3iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform3iv($1, $2)", location, v)
}

// Uniform4f fn
func (*WebGLRenderingContext) Uniform4f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.uniform4f($1, $2, $3, $4, $5)", location, x, y, z, w)
}

// Uniform4fv fn
func (*WebGLRenderingContext) Uniform4fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform4fv($1, $2)", location, v)
}

// Uniform4i fn
func (*WebGLRenderingContext) Uniform4i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int, w int) {
	js.Rewrite("$<.uniform4i($1, $2, $3, $4, $5)", location, x, y, z, w)
}

// Uniform4iv fn
func (*WebGLRenderingContext) Uniform4iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform4iv($1, $2)", location, v)
}

// UniformMatrix2fv fn
func (*WebGLRenderingContext) UniformMatrix2fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix2fv($1, $2, $3)", location, transpose, value)
}

// UniformMatrix3fv fn
func (*WebGLRenderingContext) UniformMatrix3fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix3fv($1, $2, $3)", location, transpose, value)
}

// UniformMatrix4fv fn
func (*WebGLRenderingContext) UniformMatrix4fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix4fv($1, $2, $3)", location, transpose, value)
}

// UseProgram fn
func (*WebGLRenderingContext) UseProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.useProgram($1)", program)
}

// ValidateProgram fn
func (*WebGLRenderingContext) ValidateProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.validateProgram($1)", program)
}

// VertexAttrib1f fn
func (*WebGLRenderingContext) VertexAttrib1f(indx uint, x float32) {
	js.Rewrite("$<.vertexAttrib1f($1, $2)", indx, x)
}

// VertexAttrib1fv fn
func (*WebGLRenderingContext) VertexAttrib1fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib1fv($1, $2)", indx, values)
}

// VertexAttrib2f fn
func (*WebGLRenderingContext) VertexAttrib2f(indx uint, x float32, y float32) {
	js.Rewrite("$<.vertexAttrib2f($1, $2, $3)", indx, x, y)
}

// VertexAttrib2fv fn
func (*WebGLRenderingContext) VertexAttrib2fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib2fv($1, $2)", indx, values)
}

// VertexAttrib3f fn
func (*WebGLRenderingContext) VertexAttrib3f(indx uint, x float32, y float32, z float32) {
	js.Rewrite("$<.vertexAttrib3f($1, $2, $3, $4)", indx, x, y, z)
}

// VertexAttrib3fv fn
func (*WebGLRenderingContext) VertexAttrib3fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib3fv($1, $2)", indx, values)
}

// VertexAttrib4f fn
func (*WebGLRenderingContext) VertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.vertexAttrib4f($1, $2, $3, $4, $5)", indx, x, y, z, w)
}

// VertexAttrib4fv fn
func (*WebGLRenderingContext) VertexAttrib4fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib4fv($1, $2)", indx, values)
}

// VertexAttribPointer fn
func (*WebGLRenderingContext) VertexAttribPointer(indx uint, size int, kind uint, normalized bool, stride int, offset int64) {
	js.Rewrite("$<.vertexAttribPointer($1, $2, $3, $4, $5, $6)", indx, size, kind, normalized, stride, offset)
}

// Viewport fn
func (*WebGLRenderingContext) Viewport(x int, y int, width int, height int) {
	js.Rewrite("$<.viewport($1, $2, $3, $4)", x, y, width, height)
}

// Canvas prop
func (*WebGLRenderingContext) Canvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	js.Rewrite("$<.canvas")
	return canvas
}

// DrawingBufferHeight prop
func (*WebGLRenderingContext) DrawingBufferHeight() (drawingBufferHeight int) {
	js.Rewrite("$<.drawingBufferHeight")
	return drawingBufferHeight
}

// DrawingBufferWidth prop
func (*WebGLRenderingContext) DrawingBufferWidth() (drawingBufferWidth int) {
	js.Rewrite("$<.drawingBufferWidth")
	return drawingBufferWidth
}
