package doml2deprecatedcolorproperty

import "github.com/matthewmueller/golly/js"

type DOML2DeprecatedColorProperty struct {
}

func (*DOML2DeprecatedColorProperty) GetColor() (color string) {
	js.Rewrite("$<.color")
	return color
}

func (*DOML2DeprecatedColorProperty) SetColor(color string) {
	js.Rewrite("$<.color = $1", color)
}
