package angleinstancedarrays

import "github.com/matthewmueller/golly/js"

// ANGLEInstancedArrays struct
// js:"ANGLEInstancedArrays,omit"
type ANGLEInstancedArrays struct {
}

// DrawArraysInstancedANGLE fn
// js:"drawArraysInstancedANGLE"
func (*ANGLEInstancedArrays) DrawArraysInstancedANGLE(mode uint, first int, count int, primcount int) {
	js.Rewrite("$_.drawArraysInstancedANGLE($1, $2, $3, $4)", mode, first, count, primcount)
}

// DrawElementsInstancedANGLE fn
// js:"drawElementsInstancedANGLE"
func (*ANGLEInstancedArrays) DrawElementsInstancedANGLE(mode uint, count int, kind uint, offset int64, primcount int) {
	js.Rewrite("$_.drawElementsInstancedANGLE($1, $2, $3, $4, $5)", mode, count, kind, offset, primcount)
}

// VertexAttribDivisorANGLE fn
// js:"vertexAttribDivisorANGLE"
func (*ANGLEInstancedArrays) VertexAttribDivisorANGLE(index uint, divisor uint) {
	js.Rewrite("$_.vertexAttribDivisorANGLE($1, $2)", index, divisor)
}
