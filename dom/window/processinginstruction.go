package window

import "github.com/matthewmueller/golly/js"

// ProcessingInstruction struct
// js:"ProcessingInstruction,omit"
type ProcessingInstruction struct {
	CharacterData
}

// Target prop
func (*ProcessingInstruction) Target() (target string) {
	js.Rewrite("$<.target")
	return target
}
