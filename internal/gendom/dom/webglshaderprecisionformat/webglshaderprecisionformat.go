package webglshaderprecisionformat

import "github.com/matthewmueller/golly/js"

type WebGLShaderPrecisionFormat struct {
}

func (*WebGLShaderPrecisionFormat) GetPrecision() (precision int) {
	js.Rewrite("$<.precision")
	return precision
}

func (*WebGLShaderPrecisionFormat) GetRangeMax() (rangeMax int) {
	js.Rewrite("$<.rangeMax")
	return rangeMax
}

func (*WebGLShaderPrecisionFormat) GetRangeMin() (rangeMin int) {
	js.Rewrite("$<.rangeMin")
	return rangeMin
}
