package audiolistener

import "github.com/matthewmueller/golly/js"

type AudioListener struct {
}

func (*AudioListener) SetOrientation(x float32, y float32, z float32, xUp float32, yUp float32, zUp float32) {
	js.Rewrite("$<.setOrientation($1, $2, $3, $4, $5, $6)", x, y, z, xUp, yUp, zUp)
}

func (*AudioListener) SetPosition(x float32, y float32, z float32) {
	js.Rewrite("$<.setPosition($1, $2, $3)", x, y, z)
}

func (*AudioListener) SetVelocity(x float32, y float32, z float32) {
	js.Rewrite("$<.setVelocity($1, $2, $3)", x, y, z)
}

func (*AudioListener) GetDopplerFactor() (dopplerFactor float32) {
	js.Rewrite("$<.dopplerFactor")
	return dopplerFactor
}

func (*AudioListener) SetDopplerFactor(dopplerFactor float32) {
	js.Rewrite("$<.dopplerFactor = $1", dopplerFactor)
}

func (*AudioListener) GetSpeedOfSound() (speedOfSound float32) {
	js.Rewrite("$<.speedOfSound")
	return speedOfSound
}

func (*AudioListener) SetSpeedOfSound(speedOfSound float32) {
	js.Rewrite("$<.speedOfSound = $1", speedOfSound)
}
