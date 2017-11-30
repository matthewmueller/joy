package performancenavigation

import "github.com/matthewmueller/golly/js"

// PerformanceNavigation struct
// js:"PerformanceNavigation,omit"
type PerformanceNavigation struct {
}

// ToJSON fn
// js:"toJSON"
func (*PerformanceNavigation) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// RedirectCount prop
// js:"redirectCount"
func (*PerformanceNavigation) RedirectCount() (redirectCount uint) {
	js.Rewrite("$_.redirectCount")
	return redirectCount
}

// Type prop
// js:"type"
func (*PerformanceNavigation) Type() (kind uint) {
	js.Rewrite("$_.type")
	return kind
}
