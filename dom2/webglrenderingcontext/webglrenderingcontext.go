package webglrenderingcontext

import (
	"github.com/matthewmueller/golly/dom2/htmlcanvaselement"
	"github.com/matthewmueller/golly/dom2/imagedata"
	"github.com/matthewmueller/golly/dom2/webglactiveinfo"
	"github.com/matthewmueller/golly/dom2/webglbuffer"
	"github.com/matthewmueller/golly/dom2/webglprogram"
	"github.com/matthewmueller/golly/dom2/webgluniformlocation"
	"github.com/matthewmueller/golly/js"
)

// js:"WebGLRenderingContext,omit"
type WebGLRenderingContext struct {
}

// ActiveTexture
func (*WebGLRenderingContext) ActiveTexture(texture uint) {
	js.Rewrite("$<.ActiveTexture($1)", texture)
}

// AttachShader
func (*WebGLRenderingContext) AttachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	js.Rewrite("$<.AttachShader($1, $2)", program, shader)
}

// BindAttribLocation
func (*WebGLRenderingContext) BindAttribLocation(program *webglprogram.WebGLProgram, index uint, name string) {
	js.Rewrite("$<.BindAttribLocation($1, $2, $3)", program, index, name)
}

// BindBuffer
func (*WebGLRenderingContext) BindBuffer(target uint, buffer *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.BindBuffer($1, $2)", target, buffer)
}

// BindFramebuffer
func (*WebGLRenderingContext) BindFramebuffer(target uint, framebuffer *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.BindFramebuffer($1, $2)", target, framebuffer)
}

// BindRenderbuffer
func (*WebGLRenderingContext) BindRenderbuffer(target uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.BindRenderbuffer($1, $2)", target, renderbuffer)
}

// BindTexture
func (*WebGLRenderingContext) BindTexture(target uint, texture *webgltexture.WebGLTexture) {
	js.Rewrite("$<.BindTexture($1, $2)", target, texture)
}

// BlendColor
func (*WebGLRenderingContext) BlendColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.BlendColor($1, $2, $3, $4)", red, green, blue, alpha)
}

// BlendEquation
func (*WebGLRenderingContext) BlendEquation(mode uint) {
	js.Rewrite("$<.BlendEquation($1)", mode)
}

// BlendEquationSeparate
func (*WebGLRenderingContext) BlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	js.Rewrite("$<.BlendEquationSeparate($1, $2)", modeRGB, modeAlpha)
}

// BlendFunc
func (*WebGLRenderingContext) BlendFunc(sfactor uint, dfactor uint) {
	js.Rewrite("$<.BlendFunc($1, $2)", sfactor, dfactor)
}

// BlendFuncSeparate
func (*WebGLRenderingContext) BlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	js.Rewrite("$<.BlendFuncSeparate($1, $2, $3, $4)", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// BufferData
func (*WebGLRenderingContext) BufferData(target uint, size interface{}, usage uint) {
	js.Rewrite("$<.BufferData($1, $2, $3)", target, size, usage)
}

// BufferSubData
func (*WebGLRenderingContext) BufferSubData(target uint, offset int64, data interface{}) {
	js.Rewrite("$<.BufferSubData($1, $2, $3)", target, offset, data)
}

// CheckFramebufferStatus
func (*WebGLRenderingContext) CheckFramebufferStatus(target uint) (u uint) {
	js.Rewrite("$<.CheckFramebufferStatus($1)", target)
	return u
}

// Clear
func (*WebGLRenderingContext) Clear(mask uint) {
	js.Rewrite("$<.Clear($1)", mask)
}

// ClearColor
func (*WebGLRenderingContext) ClearColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.ClearColor($1, $2, $3, $4)", red, green, blue, alpha)
}

// ClearDepth
func (*WebGLRenderingContext) ClearDepth(depth float32) {
	js.Rewrite("$<.ClearDepth($1)", depth)
}

// ClearStencil
func (*WebGLRenderingContext) ClearStencil(s int) {
	js.Rewrite("$<.ClearStencil($1)", s)
}

// ColorMask
func (*WebGLRenderingContext) ColorMask(red bool, green bool, blue bool, alpha bool) {
	js.Rewrite("$<.ColorMask($1, $2, $3, $4)", red, green, blue, alpha)
}

// CompileShader
func (*WebGLRenderingContext) CompileShader(shader *webglshader.WebGLShader) {
	js.Rewrite("$<.CompileShader($1)", shader)
}

// CompressedTexImage2d
func (*WebGLRenderingContext) CompressedTexImage2d(target uint, level int, internalformat uint, width int, height int, border int, data []byte) {
	js.Rewrite("$<.CompressedTexImage2d($1, $2, $3, $4, $5, $6, $7)", target, level, internalformat, width, height, border, data)
}

// CompressedTexSubImage2d
func (*WebGLRenderingContext) CompressedTexSubImage2d(target uint, level int, xoffset int, yoffset int, width int, height int, format uint, data []byte) {
	js.Rewrite("$<.CompressedTexSubImage2d($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, width, height, format, data)
}

// CopyTexImage2d
func (*WebGLRenderingContext) CopyTexImage2d(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	js.Rewrite("$<.CopyTexImage2d($1, $2, $3, $4, $5, $6, $7, $8)", target, level, internalformat, x, y, width, height, border)
}

// CopyTexSubImage2d
func (*WebGLRenderingContext) CopyTexSubImage2d(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	js.Rewrite("$<.CopyTexSubImage2d($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, x, y, width, height)
}

// CreateBuffer
func (*WebGLRenderingContext) CreateBuffer() (w *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.CreateBuffer()")
	return w
}

// CreateFramebuffer
func (*WebGLRenderingContext) CreateFramebuffer() (w *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.CreateFramebuffer()")
	return w
}

// CreateProgram
func (*WebGLRenderingContext) CreateProgram() (w *webglprogram.WebGLProgram) {
	js.Rewrite("$<.CreateProgram()")
	return w
}

// CreateRenderbuffer
func (*WebGLRenderingContext) CreateRenderbuffer() (w *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.CreateRenderbuffer()")
	return w
}

// CreateShader
func (*WebGLRenderingContext) CreateShader(kind uint) (w *webglshader.WebGLShader) {
	js.Rewrite("$<.CreateShader($1)", kind)
	return w
}

// CreateTexture
func (*WebGLRenderingContext) CreateTexture() (w *webgltexture.WebGLTexture) {
	js.Rewrite("$<.CreateTexture()")
	return w
}

// CullFace
func (*WebGLRenderingContext) CullFace(mode uint) {
	js.Rewrite("$<.CullFace($1)", mode)
}

// DeleteBuffer
func (*WebGLRenderingContext) DeleteBuffer(buffer *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.DeleteBuffer($1)", buffer)
}

// DeleteFramebuffer
func (*WebGLRenderingContext) DeleteFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.DeleteFramebuffer($1)", framebuffer)
}

// DeleteProgram
func (*WebGLRenderingContext) DeleteProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.DeleteProgram($1)", program)
}

// DeleteRenderbuffer
func (*WebGLRenderingContext) DeleteRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.DeleteRenderbuffer($1)", renderbuffer)
}

// DeleteShader
func (*WebGLRenderingContext) DeleteShader(shader *webglshader.WebGLShader) {
	js.Rewrite("$<.DeleteShader($1)", shader)
}

// DeleteTexture
func (*WebGLRenderingContext) DeleteTexture(texture *webgltexture.WebGLTexture) {
	js.Rewrite("$<.DeleteTexture($1)", texture)
}

// DepthFunc
func (*WebGLRenderingContext) DepthFunc(fn uint) {
	js.Rewrite("$<.DepthFunc($1)", fn)
}

// DepthMask
func (*WebGLRenderingContext) DepthMask(flag bool) {
	js.Rewrite("$<.DepthMask($1)", flag)
}

// DepthRange
func (*WebGLRenderingContext) DepthRange(zNear float32, zFar float32) {
	js.Rewrite("$<.DepthRange($1, $2)", zNear, zFar)
}

// DetachShader
func (*WebGLRenderingContext) DetachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	js.Rewrite("$<.DetachShader($1, $2)", program, shader)
}

// Disable
func (*WebGLRenderingContext) Disable(capacity uint) {
	js.Rewrite("$<.Disable($1)", capacity)
}

// DisableVertexAttribArray
func (*WebGLRenderingContext) DisableVertexAttribArray(index uint) {
	js.Rewrite("$<.DisableVertexAttribArray($1)", index)
}

// DrawArrays
func (*WebGLRenderingContext) DrawArrays(mode uint, first int, count int) {
	js.Rewrite("$<.DrawArrays($1, $2, $3)", mode, first, count)
}

// DrawElements
func (*WebGLRenderingContext) DrawElements(mode uint, count int, kind uint, offset int64) {
	js.Rewrite("$<.DrawElements($1, $2, $3, $4)", mode, count, kind, offset)
}

// Enable
func (*WebGLRenderingContext) Enable(capacity uint) {
	js.Rewrite("$<.Enable($1)", capacity)
}

// EnableVertexAttribArray
func (*WebGLRenderingContext) EnableVertexAttribArray(index uint) {
	js.Rewrite("$<.EnableVertexAttribArray($1)", index)
}

// Finish
func (*WebGLRenderingContext) Finish() {
	js.Rewrite("$<.Finish()")
}

// Flush
func (*WebGLRenderingContext) Flush() {
	js.Rewrite("$<.Flush()")
}

// FramebufferRenderbuffer
func (*WebGLRenderingContext) FramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.FramebufferRenderbuffer($1, $2, $3, $4)", target, attachment, renderbuffertarget, renderbuffer)
}

// FramebufferTexture2d
func (*WebGLRenderingContext) FramebufferTexture2d(target uint, attachment uint, textarget uint, texture *webgltexture.WebGLTexture, level int) {
	js.Rewrite("$<.FramebufferTexture2d($1, $2, $3, $4, $5)", target, attachment, textarget, texture, level)
}

// FrontFace
func (*WebGLRenderingContext) FrontFace(mode uint) {
	js.Rewrite("$<.FrontFace($1)", mode)
}

// GenerateMipmap
func (*WebGLRenderingContext) GenerateMipmap(target uint) {
	js.Rewrite("$<.GenerateMipmap($1)", target)
}

// GetActiveAttrib
func (*WebGLRenderingContext) GetActiveAttrib(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	js.Rewrite("$<.GetActiveAttrib($1, $2)", program, index)
	return w
}

// GetActiveUniform
func (*WebGLRenderingContext) GetActiveUniform(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	js.Rewrite("$<.GetActiveUniform($1, $2)", program, index)
	return w
}

// GetAttachedShaders
func (*WebGLRenderingContext) GetAttachedShaders(program *webglprogram.WebGLProgram) (w []*webglshader.WebGLShader) {
	js.Rewrite("$<.GetAttachedShaders($1)", program)
	return w
}

// GetAttribLocation
func (*WebGLRenderingContext) GetAttribLocation(program *webglprogram.WebGLProgram, name string) (i int) {
	js.Rewrite("$<.GetAttribLocation($1, $2)", program, name)
	return i
}

// GetBufferParameter
func (*WebGLRenderingContext) GetBufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.GetBufferParameter($1, $2)", target, pname)
	return i
}

// GetContextAttributes
func (*WebGLRenderingContext) GetContextAttributes() (w *webglcontextattributes.WebGLContextAttributes) {
	js.Rewrite("$<.GetContextAttributes()")
	return w
}

// GetError
func (*WebGLRenderingContext) GetError() (u uint) {
	js.Rewrite("$<.GetError()")
	return u
}

// GetExtension
func (*WebGLRenderingContext) GetExtension(name string) (i interface{}) {
	js.Rewrite("$<.GetExtension($1)", name)
	return i
}

// GetFramebufferAttachmentParameter
func (*WebGLRenderingContext) GetFramebufferAttachmentParameter(target uint, attachment uint, pname uint) (i interface{}) {
	js.Rewrite("$<.GetFramebufferAttachmentParameter($1, $2, $3)", target, attachment, pname)
	return i
}

// GetParameter
func (*WebGLRenderingContext) GetParameter(pname uint) (i interface{}) {
	js.Rewrite("$<.GetParameter($1)", pname)
	return i
}

// GetProgramInfoLog
func (*WebGLRenderingContext) GetProgramInfoLog(program *webglprogram.WebGLProgram) (s string) {
	js.Rewrite("$<.GetProgramInfoLog($1)", program)
	return s
}

// GetProgramParameter
func (*WebGLRenderingContext) GetProgramParameter(program *webglprogram.WebGLProgram, pname uint) (i interface{}) {
	js.Rewrite("$<.GetProgramParameter($1, $2)", program, pname)
	return i
}

// GetRenderbufferParameter
func (*WebGLRenderingContext) GetRenderbufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.GetRenderbufferParameter($1, $2)", target, pname)
	return i
}

// GetShaderInfoLog
func (*WebGLRenderingContext) GetShaderInfoLog(shader *webglshader.WebGLShader) (s string) {
	js.Rewrite("$<.GetShaderInfoLog($1)", shader)
	return s
}

// GetShaderParameter
func (*WebGLRenderingContext) GetShaderParameter(shader *webglshader.WebGLShader, pname uint) (i interface{}) {
	js.Rewrite("$<.GetShaderParameter($1, $2)", shader, pname)
	return i
}

// GetShaderPrecisionFormat
func (*WebGLRenderingContext) GetShaderPrecisionFormat(shadertype uint, precisiontype uint) (w *webglshaderprecisionformat.WebGLShaderPrecisionFormat) {
	js.Rewrite("$<.GetShaderPrecisionFormat($1, $2)", shadertype, precisiontype)
	return w
}

// GetShaderSource
func (*WebGLRenderingContext) GetShaderSource(shader *webglshader.WebGLShader) (s string) {
	js.Rewrite("$<.GetShaderSource($1)", shader)
	return s
}

// GetSupportedExtensions
func (*WebGLRenderingContext) GetSupportedExtensions() (s []string) {
	js.Rewrite("$<.GetSupportedExtensions()")
	return s
}

// GetTexParameter
func (*WebGLRenderingContext) GetTexParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.GetTexParameter($1, $2)", target, pname)
	return i
}

// GetUniform
func (*WebGLRenderingContext) GetUniform(program *webglprogram.WebGLProgram, location *webgluniformlocation.WebGLUniformLocation) (i interface{}) {
	js.Rewrite("$<.GetUniform($1, $2)", program, location)
	return i
}

// GetUniformLocation
func (*WebGLRenderingContext) GetUniformLocation(program *webglprogram.WebGLProgram, name string) (w *webgluniformlocation.WebGLUniformLocation) {
	js.Rewrite("$<.GetUniformLocation($1, $2)", program, name)
	return w
}

// GetVertexAttrib
func (*WebGLRenderingContext) GetVertexAttrib(index uint, pname uint) (i interface{}) {
	js.Rewrite("$<.GetVertexAttrib($1, $2)", index, pname)
	return i
}

// GetVertexAttribOffset
func (*WebGLRenderingContext) GetVertexAttribOffset(index uint, pname uint) (i int64) {
	js.Rewrite("$<.GetVertexAttribOffset($1, $2)", index, pname)
	return i
}

// Hint
func (*WebGLRenderingContext) Hint(target uint, mode uint) {
	js.Rewrite("$<.Hint($1, $2)", target, mode)
}

// IsBuffer
func (*WebGLRenderingContext) IsBuffer(buffer *webglbuffer.WebGLBuffer) (b bool) {
	js.Rewrite("$<.IsBuffer($1)", buffer)
	return b
}

// IsContextLost
func (*WebGLRenderingContext) IsContextLost() (b bool) {
	js.Rewrite("$<.IsContextLost()")
	return b
}

// IsEnabled
func (*WebGLRenderingContext) IsEnabled(capacity uint) (b bool) {
	js.Rewrite("$<.IsEnabled($1)", capacity)
	return b
}

// IsFramebuffer
func (*WebGLRenderingContext) IsFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) (b bool) {
	js.Rewrite("$<.IsFramebuffer($1)", framebuffer)
	return b
}

// IsProgram
func (*WebGLRenderingContext) IsProgram(program *webglprogram.WebGLProgram) (b bool) {
	js.Rewrite("$<.IsProgram($1)", program)
	return b
}

// IsRenderbuffer
func (*WebGLRenderingContext) IsRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) (b bool) {
	js.Rewrite("$<.IsRenderbuffer($1)", renderbuffer)
	return b
}

// IsShader
func (*WebGLRenderingContext) IsShader(shader *webglshader.WebGLShader) (b bool) {
	js.Rewrite("$<.IsShader($1)", shader)
	return b
}

// IsTexture
func (*WebGLRenderingContext) IsTexture(texture *webgltexture.WebGLTexture) (b bool) {
	js.Rewrite("$<.IsTexture($1)", texture)
	return b
}

// LineWidth
func (*WebGLRenderingContext) LineWidth(width float32) {
	js.Rewrite("$<.LineWidth($1)", width)
}

// LinkProgram
func (*WebGLRenderingContext) LinkProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.LinkProgram($1)", program)
}

// PixelStorei
func (*WebGLRenderingContext) PixelStorei(pname uint, param int) {
	js.Rewrite("$<.PixelStorei($1, $2)", pname, param)
}

// PolygonOffset
func (*WebGLRenderingContext) PolygonOffset(factor float32, units float32) {
	js.Rewrite("$<.PolygonOffset($1, $2)", factor, units)
}

// ReadPixels
func (*WebGLRenderingContext) ReadPixels(x int, y int, width int, height int, format uint, kind uint, pixels []byte) {
	js.Rewrite("$<.ReadPixels($1, $2, $3, $4, $5, $6, $7)", x, y, width, height, format, kind, pixels)
}

// RenderbufferStorage
func (*WebGLRenderingContext) RenderbufferStorage(target uint, internalformat uint, width int, height int) {
	js.Rewrite("$<.RenderbufferStorage($1, $2, $3, $4)", target, internalformat, width, height)
}

// SampleCoverage
func (*WebGLRenderingContext) SampleCoverage(value float32, invert bool) {
	js.Rewrite("$<.SampleCoverage($1, $2)", value, invert)
}

// Scissor
func (*WebGLRenderingContext) Scissor(x int, y int, width int, height int) {
	js.Rewrite("$<.Scissor($1, $2, $3, $4)", x, y, width, height)
}

// ShaderSource
func (*WebGLRenderingContext) ShaderSource(shader *webglshader.WebGLShader, source string) {
	js.Rewrite("$<.ShaderSource($1, $2)", shader, source)
}

// StencilFunc
func (*WebGLRenderingContext) StencilFunc(fn uint, ref int, mask uint) {
	js.Rewrite("$<.StencilFunc($1, $2, $3)", fn, ref, mask)
}

// StencilFuncSeparate
func (*WebGLRenderingContext) StencilFuncSeparate(face uint, fn uint, ref int, mask uint) {
	js.Rewrite("$<.StencilFuncSeparate($1, $2, $3, $4)", face, fn, ref, mask)
}

// StencilMask
func (*WebGLRenderingContext) StencilMask(mask uint) {
	js.Rewrite("$<.StencilMask($1)", mask)
}

// StencilMaskSeparate
func (*WebGLRenderingContext) StencilMaskSeparate(face uint, mask uint) {
	js.Rewrite("$<.StencilMaskSeparate($1, $2)", face, mask)
}

// StencilOp
func (*WebGLRenderingContext) StencilOp(fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.StencilOp($1, $2, $3)", fail, zfail, zpass)
}

// StencilOpSeparate
func (*WebGLRenderingContext) StencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.StencilOpSeparate($1, $2, $3, $4)", face, fail, zfail, zpass)
}

// TexImage2d
func (*WebGLRenderingContext) TexImage2d(target uint, level int, internalformat uint, format uint, kind uint, pixels *imagedata.ImageData) {
	js.Rewrite("$<.TexImage2d($1, $2, $3, $4, $5, $6)", target, level, internalformat, format, kind, pixels)
}

// TexParameterf
func (*WebGLRenderingContext) TexParameterf(target uint, pname uint, param float32) {
	js.Rewrite("$<.TexParameterf($1, $2, $3)", target, pname, param)
}

// TexParameteri
func (*WebGLRenderingContext) TexParameteri(target uint, pname uint, param int) {
	js.Rewrite("$<.TexParameteri($1, $2, $3)", target, pname, param)
}

// TexSubImage2d
func (*WebGLRenderingContext) TexSubImage2d(target uint, level int, xoffset int, yoffset int, format uint, kind uint, pixels *imagedata.ImageData) {
	js.Rewrite("$<.TexSubImage2d($1, $2, $3, $4, $5, $6, $7)", target, level, xoffset, yoffset, format, kind, pixels)
}

// Uniform1f
func (*WebGLRenderingContext) Uniform1f(location *webgluniformlocation.WebGLUniformLocation, x float32) {
	js.Rewrite("$<.Uniform1f($1, $2)", location, x)
}

// Uniform1fv
func (*WebGLRenderingContext) Uniform1fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.Uniform1fv($1, $2)", location, v)
}

// Uniform1i
func (*WebGLRenderingContext) Uniform1i(location *webgluniformlocation.WebGLUniformLocation, x int) {
	js.Rewrite("$<.Uniform1i($1, $2)", location, x)
}

// Uniform1iv
func (*WebGLRenderingContext) Uniform1iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.Uniform1iv($1, $2)", location, v)
}

// Uniform2f
func (*WebGLRenderingContext) Uniform2f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32) {
	js.Rewrite("$<.Uniform2f($1, $2, $3)", location, x, y)
}

// Uniform2fv
func (*WebGLRenderingContext) Uniform2fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.Uniform2fv($1, $2)", location, v)
}

// Uniform2i
func (*WebGLRenderingContext) Uniform2i(location *webgluniformlocation.WebGLUniformLocation, x int, y int) {
	js.Rewrite("$<.Uniform2i($1, $2, $3)", location, x, y)
}

// Uniform2iv
func (*WebGLRenderingContext) Uniform2iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.Uniform2iv($1, $2)", location, v)
}

// Uniform3f
func (*WebGLRenderingContext) Uniform3f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32) {
	js.Rewrite("$<.Uniform3f($1, $2, $3, $4)", location, x, y, z)
}

// Uniform3fv
func (*WebGLRenderingContext) Uniform3fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.Uniform3fv($1, $2)", location, v)
}

// Uniform3i
func (*WebGLRenderingContext) Uniform3i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int) {
	js.Rewrite("$<.Uniform3i($1, $2, $3, $4)", location, x, y, z)
}

// Uniform3iv
func (*WebGLRenderingContext) Uniform3iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.Uniform3iv($1, $2)", location, v)
}

// Uniform4f
func (*WebGLRenderingContext) Uniform4f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.Uniform4f($1, $2, $3, $4, $5)", location, x, y, z, w)
}

// Uniform4fv
func (*WebGLRenderingContext) Uniform4fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.Uniform4fv($1, $2)", location, v)
}

// Uniform4i
func (*WebGLRenderingContext) Uniform4i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int, w int) {
	js.Rewrite("$<.Uniform4i($1, $2, $3, $4, $5)", location, x, y, z, w)
}

// Uniform4iv
func (*WebGLRenderingContext) Uniform4iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.Uniform4iv($1, $2)", location, v)
}

// UniformMatrix2fv
func (*WebGLRenderingContext) UniformMatrix2fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.UniformMatrix2fv($1, $2, $3)", location, transpose, value)
}

// UniformMatrix3fv
func (*WebGLRenderingContext) UniformMatrix3fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.UniformMatrix3fv($1, $2, $3)", location, transpose, value)
}

// UniformMatrix4fv
func (*WebGLRenderingContext) UniformMatrix4fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.UniformMatrix4fv($1, $2, $3)", location, transpose, value)
}

// UseProgram
func (*WebGLRenderingContext) UseProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.UseProgram($1)", program)
}

// ValidateProgram
func (*WebGLRenderingContext) ValidateProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.ValidateProgram($1)", program)
}

// VertexAttrib1f
func (*WebGLRenderingContext) VertexAttrib1f(indx uint, x float32) {
	js.Rewrite("$<.VertexAttrib1f($1, $2)", indx, x)
}

// VertexAttrib1fv
func (*WebGLRenderingContext) VertexAttrib1fv(indx uint, values []float32) {
	js.Rewrite("$<.VertexAttrib1fv($1, $2)", indx, values)
}

// VertexAttrib2f
func (*WebGLRenderingContext) VertexAttrib2f(indx uint, x float32, y float32) {
	js.Rewrite("$<.VertexAttrib2f($1, $2, $3)", indx, x, y)
}

// VertexAttrib2fv
func (*WebGLRenderingContext) VertexAttrib2fv(indx uint, values []float32) {
	js.Rewrite("$<.VertexAttrib2fv($1, $2)", indx, values)
}

// VertexAttrib3f
func (*WebGLRenderingContext) VertexAttrib3f(indx uint, x float32, y float32, z float32) {
	js.Rewrite("$<.VertexAttrib3f($1, $2, $3, $4)", indx, x, y, z)
}

// VertexAttrib3fv
func (*WebGLRenderingContext) VertexAttrib3fv(indx uint, values []float32) {
	js.Rewrite("$<.VertexAttrib3fv($1, $2)", indx, values)
}

// VertexAttrib4f
func (*WebGLRenderingContext) VertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.VertexAttrib4f($1, $2, $3, $4, $5)", indx, x, y, z, w)
}

// VertexAttrib4fv
func (*WebGLRenderingContext) VertexAttrib4fv(indx uint, values []float32) {
	js.Rewrite("$<.VertexAttrib4fv($1, $2)", indx, values)
}

// VertexAttribPointer
func (*WebGLRenderingContext) VertexAttribPointer(indx uint, size int, kind uint, normalized bool, stride int, offset int64) {
	js.Rewrite("$<.VertexAttribPointer($1, $2, $3, $4, $5, $6)", indx, size, kind, normalized, stride, offset)
}

// Viewport
func (*WebGLRenderingContext) Viewport(x int, y int, width int, height int) {
	js.Rewrite("$<.Viewport($1, $2, $3, $4)", x, y, width, height)
}

// Canvas
func (*WebGLRenderingContext) Canvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	js.Rewrite("$<.Canvas")
	return canvas
}

// DrawingBufferHeight
func (*WebGLRenderingContext) DrawingBufferHeight() (drawingBufferHeight int) {
	js.Rewrite("$<.DrawingBufferHeight")
	return drawingBufferHeight
}

// DrawingBufferWidth
func (*WebGLRenderingContext) DrawingBufferWidth() (drawingBufferWidth int) {
	js.Rewrite("$<.DrawingBufferWidth")
	return drawingBufferWidth
}
