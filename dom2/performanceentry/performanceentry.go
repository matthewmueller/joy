package performanceentry

// js:"PerformanceEntry,omit"
type PerformanceEntry interface {

	// Duration
	Duration() (duration int)

	// EntryType
	EntryType() (entryType string)

	// Name
	Name() (name string)

	// StartTime
	StartTime() (startTime int)
}
