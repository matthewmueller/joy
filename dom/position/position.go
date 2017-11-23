package position

import (
	"github.com/matthewmueller/golly/dom2/coordinates"
	"github.com/matthewmueller/golly/js"
)

// Position struct
// js:"Position,omit"
type Position struct {
}

// Coords prop
func (*Position) Coords() (coords *coordinates.Coordinates) {
	js.Rewrite("$<.coords")
	return coords
}

// Timestamp prop
func (*Position) Timestamp() (timestamp int) {
	js.Rewrite("$<.timestamp")
	return timestamp
}
