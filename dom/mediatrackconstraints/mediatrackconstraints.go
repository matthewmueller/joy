package mediatrackconstraints

import "github.com/matthewmueller/golly/dom/mediatrackconstraintset"

type MediaTrackConstraints struct {
	*mediatrackconstraintset.MediaTrackConstraintSet

	advanced *[]*mediatrackconstraintset.MediaTrackConstraintSet
}
