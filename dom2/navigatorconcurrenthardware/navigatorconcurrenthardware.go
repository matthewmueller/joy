package navigatorconcurrenthardware

// js:"NavigatorConcurrentHardware,omit"
type NavigatorConcurrentHardware interface {

	// HardwareConcurrency
	HardwareConcurrency() (hardwareConcurrency uint64)
}
