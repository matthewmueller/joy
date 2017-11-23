package audiolistener

import "github.com/matthewmueller/golly/js"

// js:"AudioListener,omit"
type AudioListener struct {
}

// SetOrientation
func (*AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) {
	js.Rewrite("$<.SetOrientation($1, $2, $3, $4, $5, $6)", x, y, z, xUp, yUp, zUp)
}

// SetPosition
func (*AudioListener) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.SetPosition($1, $2, $3)", x, y, z)
}

// SetVelocity
func (*AudioListener) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.SetVelocity($1, $2, $3)", x, y, z)
}

// DopplerFactor
func (*AudioListener) DopplerFactor() (dopplerFactor float32) {
	js.Rewrite("$<.DopplerFactor")
	return dopplerFactor
}

// DopplerFactor
func (*AudioListener) SetDopplerFactor(dopplerFactor float32) {
	js.Rewrite("$<.DopplerFactor = $1", dopplerFactor)
}

// SpeedOfSound
func (*AudioListener) SpeedOfSound() (speedOfSound float32) {
	js.Rewrite("$<.SpeedOfSound")
	return speedOfSound
}

// SpeedOfSound
func (*AudioListener) SetSpeedOfSound(speedOfSound float32) {
	js.Rewrite("$<.SpeedOfSound = $1", speedOfSound)
}
