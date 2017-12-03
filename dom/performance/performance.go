package performance

import (
	"github.com/matthewmueller/joy/dom/performancenavigation"
	"github.com/matthewmueller/joy/dom/performancetiming"
	"github.com/matthewmueller/joy/macro"
)

// Performance struct
// js:"Performance,omit"
type Performance struct {
}

// ClearMarks fn
// js:"clearMarks"
func (*Performance) ClearMarks(markName *string) {
	macro.Rewrite("$_.clearMarks($1)", markName)
}

// ClearMeasures fn
// js:"clearMeasures"
func (*Performance) ClearMeasures(measureName *string) {
	macro.Rewrite("$_.clearMeasures($1)", measureName)
}

// ClearResourceTimings fn
// js:"clearResourceTimings"
func (*Performance) ClearResourceTimings() {
	macro.Rewrite("$_.clearResourceTimings()")
}

// GetEntries fn
// js:"getEntries"
func (*Performance) GetEntries() (i interface{}) {
	macro.Rewrite("$_.getEntries()")
	return i
}

// GetEntriesByName fn
// js:"getEntriesByName"
func (*Performance) GetEntriesByName(name string, entryType *string) (i interface{}) {
	macro.Rewrite("$_.getEntriesByName($1, $2)", name, entryType)
	return i
}

// GetEntriesByType fn
// js:"getEntriesByType"
func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	macro.Rewrite("$_.getEntriesByType($1)", entryType)
	return i
}

// GetMarks fn
// js:"getMarks"
func (*Performance) GetMarks(markName *string) (i interface{}) {
	macro.Rewrite("$_.getMarks($1)", markName)
	return i
}

// GetMeasures fn
// js:"getMeasures"
func (*Performance) GetMeasures(measureName *string) (i interface{}) {
	macro.Rewrite("$_.getMeasures($1)", measureName)
	return i
}

// Mark fn
// js:"mark"
func (*Performance) Mark(markName string) {
	macro.Rewrite("$_.mark($1)", markName)
}

// Measure fn
// js:"measure"
func (*Performance) Measure(measureName string, startMarkName *string, endMarkName *string) {
	macro.Rewrite("$_.measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

// Now fn
// js:"now"
func (*Performance) Now() (i int) {
	macro.Rewrite("$_.now()")
	return i
}

// SetResourceTimingBufferSize fn
// js:"setResourceTimingBufferSize"
func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	macro.Rewrite("$_.setResourceTimingBufferSize($1)", maxSize)
}

// ToJSON fn
// js:"toJSON"
func (*Performance) ToJSON() (i interface{}) {
	macro.Rewrite("$_.toJSON()")
	return i
}

// Navigation prop
// js:"navigation"
func (*Performance) Navigation() (navigation *performancenavigation.PerformanceNavigation) {
	macro.Rewrite("$_.navigation")
	return navigation
}

// Timing prop
// js:"timing"
func (*Performance) Timing() (timing *performancetiming.PerformanceTiming) {
	macro.Rewrite("$_.timing")
	return timing
}
