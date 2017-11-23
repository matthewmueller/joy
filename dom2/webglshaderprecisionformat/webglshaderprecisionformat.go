package webglshaderprecisionformat

import "github.com/matthewmueller/golly/js"

// js:"WebGLShaderPrecisionFormat,omit"
type WebGLShaderPrecisionFormat struct {
}

// Precision
func (*WebGLShaderPrecisionFormat) Precision() (precision int) {
	js.Rewrite("$<.Precision")
	return precision
}

// RangeMax
func (*WebGLShaderPrecisionFormat) RangeMax() (rangeMax int) {
	js.Rewrite("$<.RangeMax")
	return rangeMax
}

// RangeMin
func (*WebGLShaderPrecisionFormat) RangeMin() (rangeMin int) {
	js.Rewrite("$<.RangeMin")
	return rangeMin
}
