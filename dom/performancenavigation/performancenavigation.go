package performancenavigation

import "github.com/matthewmueller/joy/macro"

// PerformanceNavigation struct
// js:"PerformanceNavigation,omit"
type PerformanceNavigation struct {
}

// ToJSON fn
// js:"toJSON"
func (*PerformanceNavigation) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

// RedirectCount prop
// js:"redirectCount"
func (*PerformanceNavigation) RedirectCount() (redirectCount uint) {
	macro.Rewrite("$_.redirectCount")
	return redirectCount
}

// Type prop
// js:"type"
func (*PerformanceNavigation) Type() (kind uint) {
	macro.Rewrite("$_.type")
	return kind
}
