package position

import "github.com/matthewmueller/golly/js"

// js:"Position,omit"
type Position struct {
}

// Coords
func (*Position) Coords() (coords *coordinates.Coordinates) {
	js.Rewrite("$<.Coords")
	return coords
}

// Timestamp
func (*Position) Timestamp() (timestamp int) {
	js.Rewrite("$<.Timestamp")
	return timestamp
}
