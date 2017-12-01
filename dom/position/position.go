package position

import (
	"github.com/matthewmueller/golly/dom/coordinates"
	"github.com/matthewmueller/golly/js"
)

// Position struct
// js:"Position,omit"
type Position struct {
}

// Coords prop
// js:"coords"
func (*Position) Coords() (coords *coordinates.Coordinates) {
	js.Rewrite("$_.coords")
	return coords
}

// Timestamp prop
// js:"timestamp"
func (*Position) Timestamp() (timestamp int) {
	js.Rewrite("$_.timestamp")
	return timestamp
}
