package idb

import "github.com/matthewmueller/golly/js"

// ProcessingInstruction struct
// js:"ProcessingInstruction,omit"
type ProcessingInstruction struct {
	CharacterData
}

// Target
func (*ProcessingInstruction) Target() (target string) {
	js.Rewrite("$<.Target")
	return target
}
