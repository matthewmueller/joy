package performancemark

import (
	"github.com/matthewmueller/golly/dom/performanceentry"
	"github.com/matthewmueller/golly/js"
)

var _ performanceentry.PerformanceEntry = (*PerformanceMark)(nil)

// PerformanceMark struct
// js:"PerformanceMark,omit"
type PerformanceMark struct {
}

// Duration prop
// js:"duration"
func (*PerformanceMark) Duration() (duration int) {
	js.Rewrite("$_.duration")
	return duration
}

// EntryType prop
// js:"entryType"
func (*PerformanceMark) EntryType() (entryType string) {
	js.Rewrite("$_.entryType")
	return entryType
}

// Name prop
// js:"name"
func (*PerformanceMark) Name() (name string) {
	js.Rewrite("$_.name")
	return name
}

// StartTime prop
// js:"startTime"
func (*PerformanceMark) StartTime() (startTime int) {
	js.Rewrite("$_.startTime")
	return startTime
}
