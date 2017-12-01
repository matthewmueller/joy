package navigatorconcurrenthardware

// NavigatorConcurrentHardware interface
// js:"NavigatorConcurrentHardware"
type NavigatorConcurrentHardware interface {

	// HardwareConcurrency prop
	// js:"hardwareConcurrency"
	// jsrewrite:"$_.hardwareConcurrency"
	HardwareConcurrency() (hardwareConcurrency uint64)
}
