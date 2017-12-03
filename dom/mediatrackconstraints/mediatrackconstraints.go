package mediatrackconstraints

import "github.com/matthewmueller/joy/dom/mediatrackconstraintset"

type MediaTrackConstraints struct {
	*mediatrackconstraintset.MediaTrackConstraintSet

	advanced *[]*mediatrackconstraintset.MediaTrackConstraintSet
}
