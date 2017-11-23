package webglshaderprecisionformat

import "github.com/matthewmueller/golly/js"

// WebGLShaderPrecisionFormat struct
// js:"WebGLShaderPrecisionFormat,omit"
type WebGLShaderPrecisionFormat struct {
}

// Precision prop
func (*WebGLShaderPrecisionFormat) Precision() (precision int) {
	js.Rewrite("$<.precision")
	return precision
}

// RangeMax prop
func (*WebGLShaderPrecisionFormat) RangeMax() (rangeMax int) {
	js.Rewrite("$<.rangeMax")
	return rangeMax
}

// RangeMin prop
func (*WebGLShaderPrecisionFormat) RangeMin() (rangeMin int) {
	js.Rewrite("$<.rangeMin")
	return rangeMin
}
