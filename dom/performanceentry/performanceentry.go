package performanceentry

// PerformanceEntry interface
// js:"PerformanceEntry"
type PerformanceEntry interface {

	// Duration prop
	// js:"duration"
	// jsrewrite:"$_.duration"
	Duration() (duration int)

	// EntryType prop
	// js:"entryType"
	// jsrewrite:"$_.entryType"
	EntryType() (entryType string)

	// Name prop
	// js:"name"
	// jsrewrite:"$_.name"
	Name() (name string)

	// StartTime prop
	// js:"startTime"
	// jsrewrite:"$_.startTime"
	StartTime() (startTime int)
}
