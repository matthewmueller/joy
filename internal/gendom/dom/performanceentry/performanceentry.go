package performanceentry

import "github.com/matthewmueller/golly/js"

type PerformanceEntry struct {
}

func (*PerformanceEntry) GetDuration() (duration int) {
	js.Rewrite("$<.duration")
	return duration
}

func (*PerformanceEntry) GetEntryType() (entryType string) {
	js.Rewrite("$<.entryType")
	return entryType
}

func (*PerformanceEntry) GetName() (name string) {
	js.Rewrite("$<.name")
	return name
}

func (*PerformanceEntry) GetStartTime() (startTime int) {
	js.Rewrite("$<.startTime")
	return startTime
}
