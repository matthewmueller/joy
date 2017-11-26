package performanceentry

import "github.com/matthewmueller/golly/js"

// js:"PerformanceEntry,omit"
type PerformanceEntry interface {

	// Duration prop
	// js:"duration,rewrite=duration"
	Duration() (duration int)

	// EntryType prop
	// js:"entryType,rewrite=entrytype"
	EntryType() (entryType string)

	// Name prop
	// js:"name,rewrite=name"
	Name() (name string)

	// StartTime prop
	// js:"startTime,rewrite=starttime"
	StartTime() (startTime int)
}

// duration prop
func duration() (duration int) {
	js.Rewrite("$<.duration")
	return duration
}

// entrytype prop
func entrytype() (entryType string) {
	js.Rewrite("$<.entryType")
	return entryType
}

// name prop
func name() (name string) {
	js.Rewrite("$<.name")
	return name
}

// starttime prop
func starttime() (startTime int) {
	js.Rewrite("$<.startTime")
	return startTime
}
