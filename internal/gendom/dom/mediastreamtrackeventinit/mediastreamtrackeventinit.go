package mediastreamtrackeventinit

import "github.com/matthewmueller/golly/internal/gendom/dom/mediastreamtrack"

type MediaStreamTrackEventInit struct {
	*EventInit

	track *mediastreamtrack.MediaStreamTrack
}
