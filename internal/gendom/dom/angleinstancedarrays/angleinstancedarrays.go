package angleinstancedarrays

import "github.com/matthewmueller/golly/js"

type ANGLE_instanced_arrays struct {
}

func (*ANGLE_instanced_arrays) DrawArraysInstancedANGLE(mode uint, first int, count int, primcount int) {
	js.Rewrite("$<.drawArraysInstancedANGLE($1, $2, $3, $4)", mode, first, count, primcount)
}

func (*ANGLE_instanced_arrays) DrawElementsInstancedANGLE(mode uint, count int, kind uint, offset int64, primcount int) {
	js.Rewrite("$<.drawElementsInstancedANGLE($1, $2, $3, $4, $5)", mode, count, kind, offset, primcount)
}

func (*ANGLE_instanced_arrays) VertexAttribDivisorANGLE(index uint, divisor uint) {
	js.Rewrite("$<.vertexAttribDivisorANGLE($1, $2)", index, divisor)
}
