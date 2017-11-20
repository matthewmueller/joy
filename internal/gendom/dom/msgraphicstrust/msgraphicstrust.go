package msgraphicstrust

import "github.com/matthewmueller/golly/js"

type MSGraphicsTrust struct {
}

func (*MSGraphicsTrust) GetConstrictionActive() (constrictionActive bool) {
	js.Rewrite("$<.constrictionActive")
	return constrictionActive
}

func (*MSGraphicsTrust) GetStatus() (status string) {
	js.Rewrite("$<.status")
	return status
}
