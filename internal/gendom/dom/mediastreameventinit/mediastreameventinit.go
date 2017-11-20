package mediastreameventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/mediastream"

type MediaStreamEventInit struct {
	*EventInit

	stream *mediastream.MediaStream
}
