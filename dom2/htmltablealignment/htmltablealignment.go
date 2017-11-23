package htmltablealignment

import "github.com/matthewmueller/golly/js"

// HTMLTableAlignment struct
// js:"HTMLTableAlignment,omit"
type HTMLTableAlignment struct {
}

// Ch
func (*HTMLTableAlignment) Ch() (ch string) {
	js.Rewrite("$<.Ch")
	return ch
}

// Ch
func (*HTMLTableAlignment) SetCh(ch string) {
	js.Rewrite("$<.Ch = $1", ch)
}

// ChOff
func (*HTMLTableAlignment) ChOff() (chOff string) {
	js.Rewrite("$<.ChOff")
	return chOff
}

// ChOff
func (*HTMLTableAlignment) SetChOff(chOff string) {
	js.Rewrite("$<.ChOff = $1", chOff)
}

// VAlign
func (*HTMLTableAlignment) VAlign() (vAlign string) {
	js.Rewrite("$<.VAlign")
	return vAlign
}

// VAlign
func (*HTMLTableAlignment) SetVAlign(vAlign string) {
	js.Rewrite("$<.VAlign = $1", vAlign)
}
