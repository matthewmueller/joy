package performance

import (
	"github.com/matthewmueller/joy/dom/performancenavigation"
	"github.com/matthewmueller/joy/dom/performancetiming"
	"github.com/matthewmueller/joy/js"
)

// Performance struct
// js:"Performance,omit"
type Performance struct {
}

// ClearMarks fn
// js:"clearMarks"
func (*Performance) ClearMarks(markName *string) {
	js.Rewrite("$_.clearMarks($1)", markName)
}

// ClearMeasures fn
// js:"clearMeasures"
func (*Performance) ClearMeasures(measureName *string) {
	js.Rewrite("$_.clearMeasures($1)", measureName)
}

// ClearResourceTimings fn
// js:"clearResourceTimings"
func (*Performance) ClearResourceTimings() {
	js.Rewrite("$_.clearResourceTimings()")
}

// GetEntries fn
// js:"getEntries"
func (*Performance) GetEntries() (i interface{}) {
	js.Rewrite("$_.getEntries()")
	return i
}

// GetEntriesByName fn
// js:"getEntriesByName"
func (*Performance) GetEntriesByName(name string, entryType *string) (i interface{}) {
	js.Rewrite("$_.getEntriesByName($1, $2)", name, entryType)
	return i
}

// GetEntriesByType fn
// js:"getEntriesByType"
func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	js.Rewrite("$_.getEntriesByType($1)", entryType)
	return i
}

// GetMarks fn
// js:"getMarks"
func (*Performance) GetMarks(markName *string) (i interface{}) {
	js.Rewrite("$_.getMarks($1)", markName)
	return i
}

// GetMeasures fn
// js:"getMeasures"
func (*Performance) GetMeasures(measureName *string) (i interface{}) {
	js.Rewrite("$_.getMeasures($1)", measureName)
	return i
}

// Mark fn
// js:"mark"
func (*Performance) Mark(markName string) {
	js.Rewrite("$_.mark($1)", markName)
}

// Measure fn
// js:"measure"
func (*Performance) Measure(measureName string, startMarkName *string, endMarkName *string) {
	js.Rewrite("$_.measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

// Now fn
// js:"now"
func (*Performance) Now() (i int) {
	js.Rewrite("$_.now()")
	return i
}

// SetResourceTimingBufferSize fn
// js:"setResourceTimingBufferSize"
func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	js.Rewrite("$_.setResourceTimingBufferSize($1)", maxSize)
}

// ToJSON fn
// js:"toJSON"
func (*Performance) ToJSON() (i interface{}) {
	js.Rewrite("$_.toJSON()")
	return i
}

// Navigation prop
// js:"navigation"
func (*Performance) Navigation() (navigation *performancenavigation.PerformanceNavigation) {
	js.Rewrite("$_.navigation")
	return navigation
}

// Timing prop
// js:"timing"
func (*Performance) Timing() (timing *performancetiming.PerformanceTiming) {
	js.Rewrite("$_.timing")
	return timing
}
