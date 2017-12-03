package performancemark

import (
	"github.com/matthewmueller/joy/dom/performanceentry"
	"github.com/matthewmueller/joy/macro"
)

var _ performanceentry.PerformanceEntry = (*PerformanceMark)(nil)

// PerformanceMark struct
// js:"PerformanceMark,omit"
type PerformanceMark struct {
}

// Duration prop
// js:"duration"
func (*PerformanceMark) Duration() (duration int) {
	macro.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceMark) EntryType() (entryType string) {
	macro.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceMark) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceMark) StartTime() (startTime int) {
	macro.Rewrite("$_.startTime")
	return startTime
}
