package webglshaderprecisionformat

import "github.com/matthewmueller/joy/js"

// WebGLShaderPrecisionFormat struct
// js:"WebGLShaderPrecisionFormat,omit"
type WebGLShaderPrecisionFormat struct {
}

// Precision prop
// js:"precision"
func (*WebGLShaderPrecisionFormat) Precision() (precision int) {
	js.Rewrite("$_.precision")
	return precision
}

// RangeMax prop
// js:"rangeMax"
func (*WebGLShaderPrecisionFormat) RangeMax() (rangeMax int) {
	js.Rewrite("$_.rangeMax")
	return rangeMax
}

// RangeMin prop
// js:"rangeMin"
func (*WebGLShaderPrecisionFormat) RangeMin() (rangeMin int) {
	js.Rewrite("$_.rangeMin")
	return rangeMin
}
