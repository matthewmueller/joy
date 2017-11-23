package navigatorconcurrenthardware

import "github.com/matthewmueller/golly/js"

// NavigatorConcurrentHardware struct
// js:"NavigatorConcurrentHardware,omit"
type NavigatorConcurrentHardware struct {
}

// HardwareConcurrency prop
func (*NavigatorConcurrentHardware) HardwareConcurrency() (hardwareConcurrency uint64) {
	js.Rewrite("$<.hardwareConcurrency")
	return hardwareConcurrency
}
