package position

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/coordinates"
	"github.com/matthewmueller/golly/js"
)

type Position struct {
}

func (*Position) GetCoords() (coords *coordinates.Coordinates) {
	js.Rewrite("$<.coords")
	return coords
}

func (*Position) GetTimestamp() (timestamp int) {
	js.Rewrite("$<.timestamp")
	return timestamp
}
