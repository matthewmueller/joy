package angleinstancedarrays

import "github.com/matthewmueller/golly/js"

// ANGLEInstancedArrays struct
// js:"ANGLEInstancedArrays,omit"
type ANGLEInstancedArrays struct {
}

// DrawArraysInstancedANGLE
func (*ANGLE_instanced_arrays) DrawArraysInstancedANGLE(mode uint, first int, count int, primcount int) {
	js.Rewrite("$<.DrawArraysInstancedANGLE($1, $2, $3, $4)", mode, first, count, primcount)
}

// DrawElementsInstancedANGLE
func (*ANGLE_instanced_arrays) DrawElementsInstancedANGLE(mode uint, count int, kind uint, offset int64, primcount int) {
	js.Rewrite("$<.DrawElementsInstancedANGLE($1, $2, $3, $4, $5)", mode, count, kind, offset, primcount)
}

// VertexAttribDivisorANGLE
func (*ANGLE_instanced_arrays) VertexAttribDivisorANGLE(index uint, divisor uint) {
	js.Rewrite("$<.VertexAttribDivisorANGLE($1, $2)", index, divisor)
}
