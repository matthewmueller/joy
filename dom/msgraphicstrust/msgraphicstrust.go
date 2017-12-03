package msgraphicstrust

import "github.com/matthewmueller/joy/macro"

// MSGraphicsTrust struct
// js:"MSGraphicsTrust,omit"
type MSGraphicsTrust struct {
}

// ConstrictionActive prop
// js:"constrictionActive"
func (*MSGraphicsTrust) ConstrictionActive() (constrictionActive bool) {
	macro.Rewrite("$_.constrictionActive")
	return constrictionActive
}

// Status prop
// js:"status"
func (*MSGraphicsTrust) Status() (status string) {
	macro.Rewrite("$_.status")
	return status
}
