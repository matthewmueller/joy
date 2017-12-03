package webglshaderprecisionformat

import "github.com/matthewmueller/joy/macro"

// WebGLShaderPrecisionFormat struct
// js:"WebGLShaderPrecisionFormat,omit"
type WebGLShaderPrecisionFormat struct {
}

// Precision prop
// js:"precision"
func (*WebGLShaderPrecisionFormat) Precision() (precision int) {
	macro.Rewrite("$_.precision")
	return precision
}

// RangeMax prop
// js:"rangeMax"
func (*WebGLShaderPrecisionFormat) RangeMax() (rangeMax int) {
	macro.Rewrite("$_.rangeMax")
	return rangeMax
}

// RangeMin prop
// js:"rangeMin"
func (*WebGLShaderPrecisionFormat) RangeMin() (rangeMin int) {
	macro.Rewrite("$_.rangeMin")
	return rangeMin
}
