package js

import (
	"github.com/matthewmueller/golly/jsast"
)

// Raw statement
func Raw(src string) {
	jsast.Parse(src)
}
