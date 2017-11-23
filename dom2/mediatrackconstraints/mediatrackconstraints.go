package mediatrackconstraints

import "github.com/matthewmueller/golly/dom2/mediatrackconstraintset"

type MediaTrackConstraints struct {
	*mediatrackconstraintset.MediaTrackConstraintSet

	advanced *[]*mediatrackconstraintset.MediaTrackConstraintSet
}
