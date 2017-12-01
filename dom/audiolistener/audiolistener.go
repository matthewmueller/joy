package audiolistener

import "github.com/matthewmueller/golly/js"

// AudioListener struct
// js:"AudioListener,omit"
type AudioListener struct {
}

// SetOrientation fn
// js:"setOrientation"
func (*AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) {
	js.Rewrite("$_.setOrientation($1, $2, $3, $4, $5, $6)", x, y, z, xUp, yUp, zUp)
}

// SetPosition fn
// js:"setPosition"
func (*AudioListener) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$_.setPosition($1, $2, $3)", x, y, z)
}

// SetVelocity fn
// js:"setVelocity"
func (*AudioListener) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$_.setVelocity($1, $2, $3)", x, y, z)
}

// DopplerFactor prop
// js:"dopplerFactor"
func (*AudioListener) DopplerFactor() (dopplerFactor float32) {
	js.Rewrite("$_.dopplerFactor")
	return dopplerFactor
}

// SetDopplerFactor prop
// js:"dopplerFactor"
func (*AudioListener) SetDopplerFactor(dopplerFactor float32) {
	js.Rewrite("$_.dopplerFactor = $1", dopplerFactor)
}

// SpeedOfSound prop
// js:"speedOfSound"
func (*AudioListener) SpeedOfSound() (speedOfSound float32) {
	js.Rewrite("$_.speedOfSound")
	return speedOfSound
}

// SetSpeedOfSound prop
// js:"speedOfSound"
func (*AudioListener) SetSpeedOfSound(speedOfSound float32) {
	js.Rewrite("$_.speedOfSound = $1", speedOfSound)
}
