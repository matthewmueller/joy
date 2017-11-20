package webglrenderingcontext

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/htmlcanvaselement"
	"github.com/matthewmueller/golly/internal/gendom/dom/imagedata"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglactiveinfo"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglbuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglcontextattributes"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglframebuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglprogram"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglrenderbuffer"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglshader"
	"github.com/matthewmueller/golly/internal/gendom/dom/webglshaderprecisionformat"
	"github.com/matthewmueller/golly/internal/gendom/dom/webgltexture"
	"github.com/matthewmueller/golly/internal/gendom/dom/webgluniformlocation"
	"github.com/matthewmueller/golly/js"
)

type WebGLRenderingContext struct {
}

func (*WebGLRenderingContext) ActiveTexture(texture uint) {
	js.Rewrite("$<.activeTexture($1)", texture)
}

func (*WebGLRenderingContext) AttachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	js.Rewrite("$<.attachShader($1, $2)", program, shader)
}

func (*WebGLRenderingContext) BindAttribLocation(program *webglprogram.WebGLProgram, index uint, name string) {
	js.Rewrite("$<.bindAttribLocation($1, $2, $3)", program, index, name)
}

func (*WebGLRenderingContext) BindBuffer(target uint, buffer *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.bindBuffer($1, $2)", target, buffer)
}

func (*WebGLRenderingContext) BindFramebuffer(target uint, framebuffer *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.bindFramebuffer($1, $2)", target, framebuffer)
}

func (*WebGLRenderingContext) BindRenderbuffer(target uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.bindRenderbuffer($1, $2)", target, renderbuffer)
}

func (*WebGLRenderingContext) BindTexture(target uint, texture *webgltexture.WebGLTexture) {
	js.Rewrite("$<.bindTexture($1, $2)", target, texture)
}

func (*WebGLRenderingContext) BlendColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.blendColor($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) BlendEquation(mode uint) {
	js.Rewrite("$<.blendEquation($1)", mode)
}

func (*WebGLRenderingContext) BlendEquationSeparate(modeRGB uint, modeAlpha uint) {
	js.Rewrite("$<.blendEquationSeparate($1, $2)", modeRGB, modeAlpha)
}

func (*WebGLRenderingContext) BlendFunc(sfactor uint, dfactor uint) {
	js.Rewrite("$<.blendFunc($1, $2)", sfactor, dfactor)
}

func (*WebGLRenderingContext) BlendFuncSeparate(srcRGB uint, dstRGB uint, srcAlpha uint, dstAlpha uint) {
	js.Rewrite("$<.blendFuncSeparate($1, $2, $3, $4)", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

func (*WebGLRenderingContext) BufferData(target uint, size interface{}, usage uint) {
	js.Rewrite("$<.bufferData($1, $2, $3)", target, size, usage)
}

func (*WebGLRenderingContext) BufferSubData(target uint, offset int64, data interface{}) {
	js.Rewrite("$<.bufferSubData($1, $2, $3)", target, offset, data)
}

func (*WebGLRenderingContext) CheckFramebufferStatus(target uint) (u uint) {
	js.Rewrite("$<.checkFramebufferStatus($1)", target)
	return u
}

func (*WebGLRenderingContext) Clear(mask uint) {
	js.Rewrite("$<.clear($1)", mask)
}

func (*WebGLRenderingContext) ClearColor(red float32, green float32, blue float32, alpha float32) {
	js.Rewrite("$<.clearColor($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) ClearDepth(depth float32) {
	js.Rewrite("$<.clearDepth($1)", depth)
}

func (*WebGLRenderingContext) ClearStencil(s int) {
	js.Rewrite("$<.clearStencil($1)", s)
}

func (*WebGLRenderingContext) ColorMask(red bool, green bool, blue bool, alpha bool) {
	js.Rewrite("$<.colorMask($1, $2, $3, $4)", red, green, blue, alpha)
}

func (*WebGLRenderingContext) CompileShader(shader *webglshader.WebGLShader) {
	js.Rewrite("$<.compileShader($1)", shader)
}

func (*WebGLRenderingContext) CompressedTexImage2d(target uint, level int, internalformat uint, width int, height int, border int, data []byte) {
	js.Rewrite("$<.compressedTexImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, internalformat, width, height, border, data)
}

func (*WebGLRenderingContext) CompressedTexSubImage2d(target uint, level int, xoffset int, yoffset int, width int, height int, format uint, data []byte) {
	js.Rewrite("$<.compressedTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, width, height, format, data)
}

func (*WebGLRenderingContext) CopyTexImage2d(target uint, level int, internalformat uint, x int, y int, width int, height int, border int) {
	js.Rewrite("$<.copyTexImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, internalformat, x, y, width, height, border)
}

func (*WebGLRenderingContext) CopyTexSubImage2d(target uint, level int, xoffset int, yoffset int, x int, y int, width int, height int) {
	js.Rewrite("$<.copyTexSubImage2D($1, $2, $3, $4, $5, $6, $7, $8)", target, level, xoffset, yoffset, x, y, width, height)
}

func (*WebGLRenderingContext) CreateBuffer() (w *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.createBuffer()")
	return w
}

func (*WebGLRenderingContext) CreateFramebuffer() (w *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.createFramebuffer()")
	return w
}

func (*WebGLRenderingContext) CreateProgram() (w *webglprogram.WebGLProgram) {
	js.Rewrite("$<.createProgram()")
	return w
}

func (*WebGLRenderingContext) CreateRenderbuffer() (w *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.createRenderbuffer()")
	return w
}

func (*WebGLRenderingContext) CreateShader(kind uint) (w *webglshader.WebGLShader) {
	js.Rewrite("$<.createShader($1)", kind)
	return w
}

func (*WebGLRenderingContext) CreateTexture() (w *webgltexture.WebGLTexture) {
	js.Rewrite("$<.createTexture()")
	return w
}

func (*WebGLRenderingContext) CullFace(mode uint) {
	js.Rewrite("$<.cullFace($1)", mode)
}

func (*WebGLRenderingContext) DeleteBuffer(buffer *webglbuffer.WebGLBuffer) {
	js.Rewrite("$<.deleteBuffer($1)", buffer)
}

func (*WebGLRenderingContext) DeleteFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) {
	js.Rewrite("$<.deleteFramebuffer($1)", framebuffer)
}

func (*WebGLRenderingContext) DeleteProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.deleteProgram($1)", program)
}

func (*WebGLRenderingContext) DeleteRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.deleteRenderbuffer($1)", renderbuffer)
}

func (*WebGLRenderingContext) DeleteShader(shader *webglshader.WebGLShader) {
	js.Rewrite("$<.deleteShader($1)", shader)
}

func (*WebGLRenderingContext) DeleteTexture(texture *webgltexture.WebGLTexture) {
	js.Rewrite("$<.deleteTexture($1)", texture)
}

func (*WebGLRenderingContext) DepthFunc(fn uint) {
	js.Rewrite("$<.depthFunc($1)", fn)
}

func (*WebGLRenderingContext) DepthMask(flag bool) {
	js.Rewrite("$<.depthMask($1)", flag)
}

func (*WebGLRenderingContext) DepthRange(zNear float32, zFar float32) {
	js.Rewrite("$<.depthRange($1, $2)", zNear, zFar)
}

func (*WebGLRenderingContext) DetachShader(program *webglprogram.WebGLProgram, shader *webglshader.WebGLShader) {
	js.Rewrite("$<.detachShader($1, $2)", program, shader)
}

func (*WebGLRenderingContext) Disable(capacity uint) {
	js.Rewrite("$<.disable($1)", capacity)
}

func (*WebGLRenderingContext) DisableVertexAttribArray(index uint) {
	js.Rewrite("$<.disableVertexAttribArray($1)", index)
}

func (*WebGLRenderingContext) DrawArrays(mode uint, first int, count int) {
	js.Rewrite("$<.drawArrays($1, $2, $3)", mode, first, count)
}

func (*WebGLRenderingContext) DrawElements(mode uint, count int, kind uint, offset int64) {
	js.Rewrite("$<.drawElements($1, $2, $3, $4)", mode, count, kind, offset)
}

func (*WebGLRenderingContext) Enable(capacity uint) {
	js.Rewrite("$<.enable($1)", capacity)
}

func (*WebGLRenderingContext) EnableVertexAttribArray(index uint) {
	js.Rewrite("$<.enableVertexAttribArray($1)", index)
}

func (*WebGLRenderingContext) Finish() {
	js.Rewrite("$<.finish()")
}

func (*WebGLRenderingContext) Flush() {
	js.Rewrite("$<.flush()")
}

func (*WebGLRenderingContext) FramebufferRenderbuffer(target uint, attachment uint, renderbuffertarget uint, renderbuffer *webglrenderbuffer.WebGLRenderbuffer) {
	js.Rewrite("$<.framebufferRenderbuffer($1, $2, $3, $4)", target, attachment, renderbuffertarget, renderbuffer)
}

func (*WebGLRenderingContext) FramebufferTexture2d(target uint, attachment uint, textarget uint, texture *webgltexture.WebGLTexture, level int) {
	js.Rewrite("$<.framebufferTexture2D($1, $2, $3, $4, $5)", target, attachment, textarget, texture, level)
}

func (*WebGLRenderingContext) FrontFace(mode uint) {
	js.Rewrite("$<.frontFace($1)", mode)
}

func (*WebGLRenderingContext) GenerateMipmap(target uint) {
	js.Rewrite("$<.generateMipmap($1)", target)
}

func (*WebGLRenderingContext) GetActiveAttrib(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	js.Rewrite("$<.getActiveAttrib($1, $2)", program, index)
	return w
}

func (*WebGLRenderingContext) GetActiveUniform(program *webglprogram.WebGLProgram, index uint) (w *webglactiveinfo.WebGLActiveInfo) {
	js.Rewrite("$<.getActiveUniform($1, $2)", program, index)
	return w
}

func (*WebGLRenderingContext) GetAttachedShaders(program *webglprogram.WebGLProgram) (w []*webglshader.WebGLShader) {
	js.Rewrite("$<.getAttachedShaders($1)", program)
	return w
}

func (*WebGLRenderingContext) GetAttribLocation(program *webglprogram.WebGLProgram, name string) (i int) {
	js.Rewrite("$<.getAttribLocation($1, $2)", program, name)
	return i
}

func (*WebGLRenderingContext) GetBufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getBufferParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetContextAttributes() (w *webglcontextattributes.WebGLContextAttributes) {
	js.Rewrite("$<.getContextAttributes()")
	return w
}

func (*WebGLRenderingContext) GetError() (u uint) {
	js.Rewrite("$<.getError()")
	return u
}

func (*WebGLRenderingContext) GetExtension(name string) (i interface{}) {
	js.Rewrite("$<.getExtension($1)", name)
	return i
}

func (*WebGLRenderingContext) GetFramebufferAttachmentParameter(target uint, attachment uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getFramebufferAttachmentParameter($1, $2, $3)", target, attachment, pname)
	return i
}

func (*WebGLRenderingContext) GetParameter(pname uint) (i interface{}) {
	js.Rewrite("$<.getParameter($1)", pname)
	return i
}

func (*WebGLRenderingContext) GetProgramInfoLog(program *webglprogram.WebGLProgram) (s string) {
	js.Rewrite("$<.getProgramInfoLog($1)", program)
	return s
}

func (*WebGLRenderingContext) GetProgramParameter(program *webglprogram.WebGLProgram, pname uint) (i interface{}) {
	js.Rewrite("$<.getProgramParameter($1, $2)", program, pname)
	return i
}

func (*WebGLRenderingContext) GetRenderbufferParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getRenderbufferParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetShaderInfoLog(shader *webglshader.WebGLShader) (s string) {
	js.Rewrite("$<.getShaderInfoLog($1)", shader)
	return s
}

func (*WebGLRenderingContext) GetShaderParameter(shader *webglshader.WebGLShader, pname uint) (i interface{}) {
	js.Rewrite("$<.getShaderParameter($1, $2)", shader, pname)
	return i
}

func (*WebGLRenderingContext) GetShaderPrecisionFormat(shadertype uint, precisiontype uint) (w *webglshaderprecisionformat.WebGLShaderPrecisionFormat) {
	js.Rewrite("$<.getShaderPrecisionFormat($1, $2)", shadertype, precisiontype)
	return w
}

func (*WebGLRenderingContext) GetShaderSource(shader *webglshader.WebGLShader) (s string) {
	js.Rewrite("$<.getShaderSource($1)", shader)
	return s
}

func (*WebGLRenderingContext) GetSupportedExtensions() (s []string) {
	js.Rewrite("$<.getSupportedExtensions()")
	return s
}

func (*WebGLRenderingContext) GetTexParameter(target uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getTexParameter($1, $2)", target, pname)
	return i
}

func (*WebGLRenderingContext) GetUniform(program *webglprogram.WebGLProgram, location *webgluniformlocation.WebGLUniformLocation) (i interface{}) {
	js.Rewrite("$<.getUniform($1, $2)", program, location)
	return i
}

func (*WebGLRenderingContext) GetUniformLocation(program *webglprogram.WebGLProgram, name string) (w *webgluniformlocation.WebGLUniformLocation) {
	js.Rewrite("$<.getUniformLocation($1, $2)", program, name)
	return w
}

func (*WebGLRenderingContext) GetVertexAttrib(index uint, pname uint) (i interface{}) {
	js.Rewrite("$<.getVertexAttrib($1, $2)", index, pname)
	return i
}

func (*WebGLRenderingContext) GetVertexAttribOffset(index uint, pname uint) (i int64) {
	js.Rewrite("$<.getVertexAttribOffset($1, $2)", index, pname)
	return i
}

func (*WebGLRenderingContext) Hint(target uint, mode uint) {
	js.Rewrite("$<.hint($1, $2)", target, mode)
}

func (*WebGLRenderingContext) IsBuffer(buffer *webglbuffer.WebGLBuffer) (b bool) {
	js.Rewrite("$<.isBuffer($1)", buffer)
	return b
}

func (*WebGLRenderingContext) IsContextLost() (b bool) {
	js.Rewrite("$<.isContextLost()")
	return b
}

func (*WebGLRenderingContext) IsEnabled(capacity uint) (b bool) {
	js.Rewrite("$<.isEnabled($1)", capacity)
	return b
}

func (*WebGLRenderingContext) IsFramebuffer(framebuffer *webglframebuffer.WebGLFramebuffer) (b bool) {
	js.Rewrite("$<.isFramebuffer($1)", framebuffer)
	return b
}

func (*WebGLRenderingContext) IsProgram(program *webglprogram.WebGLProgram) (b bool) {
	js.Rewrite("$<.isProgram($1)", program)
	return b
}

func (*WebGLRenderingContext) IsRenderbuffer(renderbuffer *webglrenderbuffer.WebGLRenderbuffer) (b bool) {
	js.Rewrite("$<.isRenderbuffer($1)", renderbuffer)
	return b
}

func (*WebGLRenderingContext) IsShader(shader *webglshader.WebGLShader) (b bool) {
	js.Rewrite("$<.isShader($1)", shader)
	return b
}

func (*WebGLRenderingContext) IsTexture(texture *webgltexture.WebGLTexture) (b bool) {
	js.Rewrite("$<.isTexture($1)", texture)
	return b
}

func (*WebGLRenderingContext) LineWidth(width float32) {
	js.Rewrite("$<.lineWidth($1)", width)
}

func (*WebGLRenderingContext) LinkProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.linkProgram($1)", program)
}

func (*WebGLRenderingContext) PixelStorei(pname uint, param int) {
	js.Rewrite("$<.pixelStorei($1, $2)", pname, param)
}

func (*WebGLRenderingContext) PolygonOffset(factor float32, units float32) {
	js.Rewrite("$<.polygonOffset($1, $2)", factor, units)
}

func (*WebGLRenderingContext) ReadPixels(x int, y int, width int, height int, format uint, kind uint, pixels []byte) {
	js.Rewrite("$<.readPixels($1, $2, $3, $4, $5, $6, $7)", x, y, width, height, format, kind, pixels)
}

func (*WebGLRenderingContext) RenderbufferStorage(target uint, internalformat uint, width int, height int) {
	js.Rewrite("$<.renderbufferStorage($1, $2, $3, $4)", target, internalformat, width, height)
}

func (*WebGLRenderingContext) SampleCoverage(value float32, invert bool) {
	js.Rewrite("$<.sampleCoverage($1, $2)", value, invert)
}

func (*WebGLRenderingContext) Scissor(x int, y int, width int, height int) {
	js.Rewrite("$<.scissor($1, $2, $3, $4)", x, y, width, height)
}

func (*WebGLRenderingContext) ShaderSource(shader *webglshader.WebGLShader, source string) {
	js.Rewrite("$<.shaderSource($1, $2)", shader, source)
}

func (*WebGLRenderingContext) StencilFunc(fn uint, ref int, mask uint) {
	js.Rewrite("$<.stencilFunc($1, $2, $3)", fn, ref, mask)
}

func (*WebGLRenderingContext) StencilFuncSeparate(face uint, fn uint, ref int, mask uint) {
	js.Rewrite("$<.stencilFuncSeparate($1, $2, $3, $4)", face, fn, ref, mask)
}

func (*WebGLRenderingContext) StencilMask(mask uint) {
	js.Rewrite("$<.stencilMask($1)", mask)
}

func (*WebGLRenderingContext) StencilMaskSeparate(face uint, mask uint) {
	js.Rewrite("$<.stencilMaskSeparate($1, $2)", face, mask)
}

func (*WebGLRenderingContext) StencilOp(fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.stencilOp($1, $2, $3)", fail, zfail, zpass)
}

func (*WebGLRenderingContext) StencilOpSeparate(face uint, fail uint, zfail uint, zpass uint) {
	js.Rewrite("$<.stencilOpSeparate($1, $2, $3, $4)", face, fail, zfail, zpass)
}

func (*WebGLRenderingContext) TexImage2d(target uint, level int, internalformat uint, format uint, kind uint, pixels *imagedata.ImageData) {
	js.Rewrite("$<.texImage2D($1, $2, $3, $4, $5, $6)", target, level, internalformat, format, kind, pixels)
}

func (*WebGLRenderingContext) TexParameterf(target uint, pname uint, param float32) {
	js.Rewrite("$<.texParameterf($1, $2, $3)", target, pname, param)
}

func (*WebGLRenderingContext) TexParameteri(target uint, pname uint, param int) {
	js.Rewrite("$<.texParameteri($1, $2, $3)", target, pname, param)
}

func (*WebGLRenderingContext) TexSubImage2d(target uint, level int, xoffset int, yoffset int, format uint, kind uint, pixels *imagedata.ImageData) {
	js.Rewrite("$<.texSubImage2D($1, $2, $3, $4, $5, $6, $7)", target, level, xoffset, yoffset, format, kind, pixels)
}

func (*WebGLRenderingContext) Uniform1f(location *webgluniformlocation.WebGLUniformLocation, x float32) {
	js.Rewrite("$<.uniform1f($1, $2)", location, x)
}

func (*WebGLRenderingContext) Uniform1fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform1fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform1i(location *webgluniformlocation.WebGLUniformLocation, x int) {
	js.Rewrite("$<.uniform1i($1, $2)", location, x)
}

func (*WebGLRenderingContext) Uniform1iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform1iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform2f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32) {
	js.Rewrite("$<.uniform2f($1, $2, $3)", location, x, y)
}

func (*WebGLRenderingContext) Uniform2fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform2fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform2i(location *webgluniformlocation.WebGLUniformLocation, x int, y int) {
	js.Rewrite("$<.uniform2i($1, $2, $3)", location, x, y)
}

func (*WebGLRenderingContext) Uniform2iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform2iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform3f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32) {
	js.Rewrite("$<.uniform3f($1, $2, $3, $4)", location, x, y, z)
}

func (*WebGLRenderingContext) Uniform3fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform3fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform3i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int) {
	js.Rewrite("$<.uniform3i($1, $2, $3, $4)", location, x, y, z)
}

func (*WebGLRenderingContext) Uniform3iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform3iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform4f(location *webgluniformlocation.WebGLUniformLocation, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.uniform4f($1, $2, $3, $4, $5)", location, x, y, z, w)
}

func (*WebGLRenderingContext) Uniform4fv(location *webgluniformlocation.WebGLUniformLocation, v []float32) {
	js.Rewrite("$<.uniform4fv($1, $2)", location, v)
}

func (*WebGLRenderingContext) Uniform4i(location *webgluniformlocation.WebGLUniformLocation, x int, y int, z int, w int) {
	js.Rewrite("$<.uniform4i($1, $2, $3, $4, $5)", location, x, y, z, w)
}

func (*WebGLRenderingContext) Uniform4iv(location *webgluniformlocation.WebGLUniformLocation, v []int32) {
	js.Rewrite("$<.uniform4iv($1, $2)", location, v)
}

func (*WebGLRenderingContext) UniformMatrix2fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix2fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UniformMatrix3fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix3fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UniformMatrix4fv(location *webgluniformlocation.WebGLUniformLocation, transpose bool, value []float32) {
	js.Rewrite("$<.uniformMatrix4fv($1, $2, $3)", location, transpose, value)
}

func (*WebGLRenderingContext) UseProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.useProgram($1)", program)
}

func (*WebGLRenderingContext) ValidateProgram(program *webglprogram.WebGLProgram) {
	js.Rewrite("$<.validateProgram($1)", program)
}

func (*WebGLRenderingContext) VertexAttrib1f(indx uint, x float32) {
	js.Rewrite("$<.vertexAttrib1f($1, $2)", indx, x)
}

func (*WebGLRenderingContext) VertexAttrib1fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib1fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib2f(indx uint, x float32, y float32) {
	js.Rewrite("$<.vertexAttrib2f($1, $2, $3)", indx, x, y)
}

func (*WebGLRenderingContext) VertexAttrib2fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib2fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib3f(indx uint, x float32, y float32, z float32) {
	js.Rewrite("$<.vertexAttrib3f($1, $2, $3, $4)", indx, x, y, z)
}

func (*WebGLRenderingContext) VertexAttrib3fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib3fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttrib4f(indx uint, x float32, y float32, z float32, w float32) {
	js.Rewrite("$<.vertexAttrib4f($1, $2, $3, $4, $5)", indx, x, y, z, w)
}

func (*WebGLRenderingContext) VertexAttrib4fv(indx uint, values []float32) {
	js.Rewrite("$<.vertexAttrib4fv($1, $2)", indx, values)
}

func (*WebGLRenderingContext) VertexAttribPointer(indx uint, size int, kind uint, normalized bool, stride int, offset int64) {
	js.Rewrite("$<.vertexAttribPointer($1, $2, $3, $4, $5, $6)", indx, size, kind, normalized, stride, offset)
}

func (*WebGLRenderingContext) Viewport(x int, y int, width int, height int) {
	js.Rewrite("$<.viewport($1, $2, $3, $4)", x, y, width, height)
}

func (*WebGLRenderingContext) GetCanvas() (canvas *htmlcanvaselement.HTMLCanvasElement) {
	js.Rewrite("$<.canvas")
	return canvas
}

func (*WebGLRenderingContext) GetDrawingBufferHeight() (drawingBufferHeight int) {
	js.Rewrite("$<.drawingBufferHeight")
	return drawingBufferHeight
}

func (*WebGLRenderingContext) GetDrawingBufferWidth() (drawingBufferWidth int) {
	js.Rewrite("$<.drawingBufferWidth")
	return drawingBufferWidth
}
