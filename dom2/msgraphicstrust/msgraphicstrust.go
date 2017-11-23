package msgraphicstrust

import "github.com/matthewmueller/golly/js"

// js:"MSGraphicsTrust,omit"
type MSGraphicsTrust struct {
}

// ConstrictionActive
func (*MSGraphicsTrust) ConstrictionActive() (constrictionActive bool) {
	js.Rewrite("$<.ConstrictionActive")
	return constrictionActive
}

// Status
func (*MSGraphicsTrust) Status() (status string) {
	js.Rewrite("$<.Status")
	return status
}
