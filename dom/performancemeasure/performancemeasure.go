package performancemeasure

import (
	"github.com/matthewmueller/golly/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

var _ performanceentry.PerformanceEntry = (*PerformanceMeasure)(nil)

// PerformanceMeasure struct
// js:"PerformanceMeasure,omit"
type PerformanceMeasure struct {
}

// Duration prop
// js:"duration"
func (*PerformanceMeasure) Duration() (duration int) {
	js.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceMeasure) EntryType() (entryType string) {
	js.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceMeasure) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceMeasure) StartTime() (startTime int) {
	js.Rewrite("$_.startTime")
	return startTime
}
