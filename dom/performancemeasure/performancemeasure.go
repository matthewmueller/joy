package performancemeasure

import (
	"github.com/matthewmueller/joy/dom/performanceentry"
	"github.com/matthewmueller/joy/macro"
)

var _ performanceentry.PerformanceEntry = (*PerformanceMeasure)(nil)

// PerformanceMeasure struct
// js:"PerformanceMeasure,omit"
type PerformanceMeasure struct {
}

// Duration prop
// js:"duration"
func (*PerformanceMeasure) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceMeasure) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceMeasure) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceMeasure) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
