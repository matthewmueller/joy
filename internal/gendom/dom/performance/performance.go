package performance

import (
	"github.com/matthewmueller/golly/internal/gendom/dom/performancenavigation"
	"github.com/matthewmueller/golly/internal/gendom/dom/performancetiming"
	"github.com/matthewmueller/golly/js"
)

type Performance struct {
}

func (*Performance) ClearMarks(markName string) {
	js.Rewrite("$<.clearMarks($1)", markName)
}

func (*Performance) ClearMeasures(measureName string) {
	js.Rewrite("$<.clearMeasures($1)", measureName)
}

func (*Performance) ClearResourceTimings() {
	js.Rewrite("$<.clearResourceTimings()")
}

func (*Performance) GetEntries() (i interface{}) {
	js.Rewrite("$<.getEntries()")
	return i
}

func (*Performance) GetEntriesByName(name string, entryType string) (i interface{}) {
	js.Rewrite("$<.getEntriesByName($1, $2)", name, entryType)
	return i
}

func (*Performance) GetEntriesByType(entryType string) (i interface{}) {
	js.Rewrite("$<.getEntriesByType($1)", entryType)
	return i
}

func (*Performance) GetMarks(markName string) (i interface{}) {
	js.Rewrite("$<.getMarks($1)", markName)
	return i
}

func (*Performance) GetMeasures(measureName string) (i interface{}) {
	js.Rewrite("$<.getMeasures($1)", measureName)
	return i
}

func (*Performance) Mark(markName string) {
	js.Rewrite("$<.mark($1)", markName)
}

func (*Performance) Measure(measureName string, startMarkName string, endMarkName string) {
	js.Rewrite("$<.measure($1, $2, $3)", measureName, startMarkName, endMarkName)
}

func (*Performance) Now() (i int) {
	js.Rewrite("$<.now()")
	return i
}

func (*Performance) SetResourceTimingBufferSize(maxSize uint) {
	js.Rewrite("$<.setResourceTimingBufferSize($1)", maxSize)
}

func (*Performance) ToJSON() (i interface{}) {
	js.Rewrite("$<.toJSON()")
	return i
}

func (*Performance) GetNavigation() (navigation *performancenavigation.PerformanceNavigation) {
	js.Rewrite("$<.navigation")
	return navigation
}

func (*Performance) GetTiming() (timing *performancetiming.PerformanceTiming) {
	js.Rewrite("$<.timing")
	return timing
}
