package audiolistener

import "github.com/matthewmueller/golly/js"

// AudioListener struct
// js:"AudioListener,omit"
type AudioListener struct {
}

// SetOrientation fn
func (*AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) {
	js.Rewrite("$<.setOrientation($1, $2, $3, $4, $5, $6)", x, y, z, xUp, yUp, zUp)
}

// SetPosition fn
func (*AudioListener) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.setPosition($1, $2, $3)", x, y, z)
}

// SetVelocity fn
func (*AudioListener) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.setVelocity($1, $2, $3)", x, y, z)
}

// DopplerFactor prop
func (*AudioListener) DopplerFactor() (dopplerFactor float32) {
	js.Rewrite("$<.dopplerFactor")
	return dopplerFactor
}

// DopplerFactor prop
func (*AudioListener) SetDopplerFactor(dopplerFactor float32) {
	js.Rewrite("$<.dopplerFactor = $1", dopplerFactor)
}

// SpeedOfSound prop
func (*AudioListener) SpeedOfSound() (speedOfSound float32) {
	js.Rewrite("$<.speedOfSound")
	return speedOfSound
}

// SpeedOfSound prop
func (*AudioListener) SetSpeedOfSound(speedOfSound float32) {
	js.Rewrite("$<.speedOfSound = $1", speedOfSound)
}
