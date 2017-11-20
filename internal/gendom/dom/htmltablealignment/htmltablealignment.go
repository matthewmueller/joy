package htmltablealignment

import "github.com/matthewmueller/golly/js"

type HTMLTableAlignment struct {
}

func (*HTMLTableAlignment) GetCh() (ch string) {
	js.Rewrite("$<.ch")
	return ch
}

func (*HTMLTableAlignment) SetCh(ch string) {
	js.Rewrite("$<.ch = $1", ch)
}

func (*HTMLTableAlignment) GetChOff() (chOff string) {
	js.Rewrite("$<.chOff")
	return chOff
}

func (*HTMLTableAlignment) SetChOff(chOff string) {
	js.Rewrite("$<.chOff = $1", chOff)
}

func (*HTMLTableAlignment) GetVAlign() (vAlign string) {
	js.Rewrite("$<.vAlign")
	return vAlign
}

func (*HTMLTableAlignment) SetVAlign(vAlign string) {
	js.Rewrite("$<.vAlign = $1", vAlign)
}
