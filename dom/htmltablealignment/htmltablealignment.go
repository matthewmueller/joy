package htmltablealignment

// HTMLTableAlignment interface
// js:"HTMLTableAlignment"
type HTMLTableAlignment interface {

	// Ch prop
	// js:"ch"
	// jsrewrite:"$_.ch"
	Ch() (ch string)

	// SetCh prop
	// js:"ch"
	// jsrewrite:"$_.ch = $1"
	SetCh(ch string)

	// ChOff prop
	// js:"chOff"
	// jsrewrite:"$_.chOff"
	ChOff() (chOff string)

	// SetChOff prop
	// js:"chOff"
	// jsrewrite:"$_.chOff = $1"
	SetChOff(chOff string)

	// VAlign prop
	// js:"vAlign"
	// jsrewrite:"$_.vAlign"
	VAlign() (vAlign string)

	// SetVAlign prop
	// js:"vAlign"
	// jsrewrite:"$_.vAlign = $1"
	SetVAlign(vAlign string)
}
