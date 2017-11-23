package msgraphicstrust

import "github.com/matthewmueller/golly/js"

// MSGraphicsTrust struct
// js:"MSGraphicsTrust,omit"
type MSGraphicsTrust struct {
}

// ConstrictionActive prop
func (*MSGraphicsTrust) ConstrictionActive() (constrictionActive bool) {
	js.Rewrite("$<.constrictionActive")
	return constrictionActive
}

// Status prop
func (*MSGraphicsTrust) Status() (status string) {
	js.Rewrite("$<.status")
	return status
}
