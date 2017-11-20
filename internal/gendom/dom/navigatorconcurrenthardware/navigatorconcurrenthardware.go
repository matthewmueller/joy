package navigatorconcurrenthardware

import "github.com/matthewmueller/golly/js"

type NavigatorConcurrentHardware struct {
}

func (*NavigatorConcurrentHardware) GetHardwareConcurrency() (hardwareConcurrency uint64) {
	js.Rewrite("$<.hardwareConcurrency")
	return hardwareConcurrency
}
