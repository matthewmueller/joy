package datacue

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/texttrackcue"
	"github.com/matthewmueller/golly/js"
)

type DataCue struct {
	*texttrackcue.TextTrackCue
}

func (*DataCue) GetData() (data []byte) {
	js.Rewrite("$<.data")
	return data
}

func (*DataCue) SetData(data []byte) {
	js.Rewrite("$<.data = $1", data)
}
