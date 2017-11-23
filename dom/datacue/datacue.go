package datacue

import (
	"github.com/matthewmueller/golly/dom/texttrack"
	"github.com/matthewmueller/golly/js"
)

// DataCue struct
// js:"DataCue,omit"
type DataCue struct {
	texttrack.TextTrackCue
}

// Data prop
func (*DataCue) Data() (data []byte) {
	js.Rewrite("$<.data")
	return data
}

// Data prop
func (*DataCue) SetData(data []byte) {
	js.Rewrite("$<.data = $1", data)
}
