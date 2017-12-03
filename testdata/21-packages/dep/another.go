package dep

import (
	"github.com/matthewmueller/joy/testdata/21-packages/deepdep"
)

// A var
var A = "A"

// B const
// const B = "B"

// C const
var C = "C"

func another() string {
	return "another"
}

// Another fn
func Another() string {
	deep := deepdep.New()
	return "exported another " + deep.String()
}

// Unused fn
func Unused() string {
	return "i'm not used"
}
