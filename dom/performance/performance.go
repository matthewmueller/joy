package performance

import (
	"github.com/matthewmueller/golly/dom2/performancenavigation"
	"github.com/matthewmueller/golly/dom2/performancetiming"
	"github.com/matthewmueller/golly/js"
)

// Performance struct
// js:"Performance,omit"
type Performance struct {
}

// ClearMarks fn
func (*Performance) ClearMarks(markName *string) {
	js.Rewrite("$<.clearMarks($1)", markName)
}

// ClearMeasures fn
func (*Performance) ClearMeasures(measureName *string) {
	js.Rewrite("$<.clearMeasures($1)", measureName)
}

// ClearResourceTimings fn
func (*Performance) ClearResourceTimings() {
	js.Rewrite("$<.clearResourceTimings()")
}

// GetEntries fn
func (*Performance) GetEntries() (i interface{}) {
	js.Rewrite("$<.getEntries()")
	return i
}

// GetEntriesByName fn
func (*Performance) GetEntriesByName(name string, entryType *string) (i interface{}) {
	js.Rewrite("$<.getEntriesByName($1, $2)", name, entryType)
	return i
}

// GetEntriesByType fn
func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	js.Rewrite("$<.getEntriesByType($1)", entryType)
	return i
}

// GetMarks fn
func (*Performance) GetMarks(markName *string) (i interface{}) {
	js.Rewrite("$<.getMarks($1)", markName)
	return i
}

// GetMeasures fn
func (*Performance) GetMeasures(measureName *string) (i interface{}) {
	js.Rewrite("$<.getMeasures($1)", measureName)
	return i
}

// Mark fn
func (*Performance) Mark(markName string) {
	js.Rewrite("$<.mark($1)", markName)
}

// Measure fn
func (*Performance) Measure(measureName string, startMarkName *string, endMarkName *string) {
	js.Rewrite("$<.measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

// Now fn
func (*Performance) Now() (i int) {
	js.Rewrite("$<.now()")
	return i
}

// SetResourceTimingBufferSize fn
func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	js.Rewrite("$<.setResourceTimingBufferSize($1)", maxSize)
}

// ToJSON fn
func (*Performance) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

// Navigation prop
func (*Performance) Navigation() (navigation *performancenavigation.PerformanceNavigation) {
	js.Rewrite("$<.navigation")
	return navigation
}

// Timing prop
func (*Performance) Timing() (timing *performancetiming.PerformanceTiming) {
	js.Rewrite("$<.timing")
	return timing
}
