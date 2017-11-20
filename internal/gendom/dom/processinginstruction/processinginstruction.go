package processinginstruction

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/characterdata"
	"github.com/matthewmueller/golly/js"
)

type ProcessingInstruction struct {
	*characterdata.CharacterData
}

func (*ProcessingInstruction) GetTarget() (target string) {
	js.Rewrite("$<.target")
	return target
}
