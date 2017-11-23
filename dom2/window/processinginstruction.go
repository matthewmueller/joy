package window

import "github.com/matthewmueller/golly/js"

// js:"ProcessingInstruction,omit"
type ProcessingInstruction struct {
	CharacterData
}

// Target
func (*ProcessingInstruction) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}
