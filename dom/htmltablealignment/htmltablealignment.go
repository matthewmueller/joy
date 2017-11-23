package htmltablealignment

import "github.com/matthewmueller/golly/js"

// HTMLTableAlignment struct
// js:"HTMLTableAlignment,omit"
type HTMLTableAlignment struct {
}

// Ch prop
func (*HTMLTableAlignment) Ch() (ch string) {
	js.Rewrite("$<.ch")
	return ch
}

// Ch prop
func (*HTMLTableAlignment) SetCh(ch string) {
	js.Rewrite("$<.ch = $1", ch)
}

// ChOff prop
func (*HTMLTableAlignment) ChOff() (chOff string) {
	js.Rewrite("$<.chOff")
	return chOff
}

// ChOff prop
func (*HTMLTableAlignment) SetChOff(chOff string) {
	js.Rewrite("$<.chOff = $1", chOff)
}

// VAlign prop
func (*HTMLTableAlignment) VAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

// VAlign prop
func (*HTMLTableAlignment) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}
