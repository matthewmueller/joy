package datacue

import (
	"github.com/matthewmueller/golly/dom2/texttrack"
	"github.com/matthewmueller/golly/js"
)

// DataCue struct
// js:"DataCue,omit"
type DataCue struct {
	texttrack.TextTrackCue
}

// Data
func (*DataCue) Data() (data []byte) {
	js.Rewrite("$<.Data")
	return data
}

// Data
func (*DataCue) SetData(data []byte) {
	js.Rewrite("$<.Data = $1", data)
}
