package navigatorconcurrenthardware

import "github.com/matthewmueller/golly/js"

// NavigatorConcurrentHardware struct
// js:"NavigatorConcurrentHardware,omit"
type NavigatorConcurrentHardware struct {
}

// HardwareConcurrency
func (*NavigatorConcurrentHardware) HardwareConcurrency() (hardwareConcurrency uint64) {
	js.Rewrite("$<.HardwareConcurrency")
	return hardwareConcurrency
}
