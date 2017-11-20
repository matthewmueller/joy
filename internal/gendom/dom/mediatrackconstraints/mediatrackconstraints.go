package mediatrackconstraints

import "github.com/matthewmueller/golly/internal/gendom/dom/mediatrackconstraintset"

type MediaTrackConstraints struct {
	*MediaTrackConstraintSet

	advanced *[]*mediatrackconstraintset.MediaTrackConstraintSet
}
