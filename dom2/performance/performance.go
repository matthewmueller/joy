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

// ClearMarks
func (*Performance) ClearMarks(markName *string) {
	js.Rewrite("$<.ClearMarks($1)", markName)
}

// ClearMeasures
func (*Performance) ClearMeasures(measureName *string) {
	js.Rewrite("$<.ClearMeasures($1)", measureName)
}

// ClearResourceTimings
func (*Performance) ClearResourceTimings() {
	js.Rewrite("$<.ClearResourceTimings()")
}

// GetEntries
func (*Performance) GetEntries() (i interface{}) {
	js.Rewrite("$<.GetEntries()")
	return i
}

// GetEntriesByName
func (*Performance) GetEntriesByName(name string, entryType *string) (i interface{}) {
	js.Rewrite("$<.GetEntriesByName($1, $2)", name, entryType)
	return i
}

// GetEntriesByType
func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	js.Rewrite("$<.GetEntriesByType($1)", entryType)
	return i
}

// GetMarks
func (*Performance) GetMarks(markName *string) (i interface{}) {
	js.Rewrite("$<.GetMarks($1)", markName)
	return i
}

// GetMeasures
func (*Performance) GetMeasures(measureName *string) (i interface{}) {
	js.Rewrite("$<.GetMeasures($1)", measureName)
	return i
}

// Mark
func (*Performance) Mark(markName string) {
	js.Rewrite("$<.Mark($1)", markName)
}

// Measure
func (*Performance) Measure(measureName string, startMarkName *string, endMarkName *string) {
	js.Rewrite("$<.Measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

// Now
func (*Performance) Now() (i int) {
	js.Rewrite("$<.Now()")
	return i
}

// SetResourceTimingBufferSize
func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	js.Rewrite("$<.SetResourceTimingBufferSize($1)", maxSize)
}

// ToJSON
func (*Performance) ToJSON() (i interface{}) {
	js.Rewrite("$<.ToJSON()")
	return i
}

// Navigation
func (*Performance) Navigation() (navigation *performancenavigation.PerformanceNavigation) {
	js.Rewrite("$<.Navigation")
	return navigation
}

// Timing
func (*Performance) Timing() (timing *performancetiming.PerformanceTiming) {
	js.Rewrite("$<.Timing")
	return timing
}
