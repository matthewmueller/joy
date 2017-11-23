package angleinstancedarrays

import "github.com/matthewmueller/golly/js"

// ANGLEInstancedArrays struct
// js:"ANGLEInstancedArrays,omit"
type ANGLEInstancedArrays struct {
}

// DrawArraysInstancedANGLE fn
func (*ANGLE_instanced_arrays) DrawArraysInstancedANGLE(mode uint, first int, count int, primcount int) {
	js.Rewrite("$<.drawArraysInstancedANGLE($1, $2, $3, $4)", mode, first, count, primcount)
}

// DrawElementsInstancedANGLE fn
func (*ANGLE_instanced_arrays) DrawElementsInstancedANGLE(mode uint, count int, kind uint, offset int64, primcount int) {
	js.Rewrite("$<.drawElementsInstancedANGLE($1, $2, $3, $4, $5)", mode, count, kind, offset, primcount)
}

// VertexAttribDivisorANGLE fn
func (*ANGLE_instanced_arrays) VertexAttribDivisorANGLE(index uint, divisor uint) {
	js.Rewrite("$<.vertexAttribDivisorANGLE($1, $2)", index, divisor)
}
