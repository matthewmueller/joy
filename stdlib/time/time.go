package time

import (
	"github.com/matthewmueller/joy/macro"
)

// Time struct
// type Time struct{}

// A Duration represents the elapsed time between two instants
// as an int64 nanosecond count. The representation limits the
// largest representable duration to approximately 290 years.
// type Duration int64

// const (
// 	minDuration Duration = -1 << 63
// 	maxDuration Duration = 1<<63 - 1
// )

// Common durations. There is no definition for units of Day or larger
// to avoid confusion across daylight savings time zone transitions.
//
// To count the number of units in a Duration, divide:
//	second := time.Second
//	fmt.Print(int64(second/time.Millisecond)) // prints 1000
//
// To convert an integer number of units to a Duration, multiply:
//	seconds := 10
//	fmt.Print(time.Duration(seconds)*time.Second) // prints 10s

// Millisecond const
var Millisecond = func() int {
	macro.Rewrite("1")
	return 1000 * 1000 * 1
}()

// Second const
var Second = func() int {
	macro.Rewrite("1000")
	return 1000 * 1000 * 1000 * 1
}()

// Minute const
var Minute = func() int {
	macro.Rewrite("60000")
	return 60 * 1000 * 1000 * 1000 * 1
}()

// Hour const
var Hour = func() int {
	macro.Rewrite("3600000")
	return 60 * 60 * 1000 * 1000 * 1000 * 1
}()

// Sleep pauses the current goroutine for at least the duration d.
// A negative or zero duration causes Sleep to return immediately.
func Sleep(d int) {
	macro.Rewrite("await new Promise(function(resolve, reject) { setTimeout(resolve, $1) })", d)
}
