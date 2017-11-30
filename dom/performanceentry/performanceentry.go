package performanceentry

// PerformanceEntry interface
// js:"PerformanceEntry"
type PerformanceEntry interface {

	// Duration prop
	// js:"duration"
	Duration() (duration int)

	// EntryType prop
	// js:"entryType"
	EntryType() (entryType string)

	// Name prop
	// js:"name"
	Name() (name string)

	// StartTime prop
	// js:"startTime"
	StartTime() (startTime int)
}
