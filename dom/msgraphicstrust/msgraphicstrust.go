package msgraphicstrust

import "github.com/matthewmueller/golly/js"

// MSGraphicsTrust struct
// js:"MSGraphicsTrust,omit"
type MSGraphicsTrust struct {
}

// ConstrictionActive prop
// js:"constrictionActive"
func (*MSGraphicsTrust) ConstrictionActive() (constrictionActive bool) {
	js.Rewrite("$_.constrictionActive")
	return constrictionActive
}

// Status prop
// js:"status"
func (*MSGraphicsTrust) Status() (status string) {
	js.Rewrite("$_.status")
	return status
}
