package perfwidgetexternal

import "github.com/matthewmueller/golly/js"

// PerfWidgetExternal struct
// js:"PerfWidgetExternal,omit"
type PerfWidgetExternal struct {
}

// AddEventListener
func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	js.Rewrite("$<.AddEventListener($1, $2)", eventType, callback)
}

// GetMemoryUsage
func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	js.Rewrite("$<.GetMemoryUsage()")
	return u
}

// GetProcessCPUUsage
func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	js.Rewrite("$<.GetProcessCPUUsage()")
	return u
}

// GetRecentCPUUsage
func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.GetRecentCPUUsage($1)", last)
	return i
}

// GetRecentFrames
func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	js.Rewrite("$<.GetRecentFrames($1)", last)
	return i
}

// GetRecentMemoryUsage
func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	js.Rewrite("$<.GetRecentMemoryUsage($1)", last)
	return i
}

// GetRecentPaintRequests
func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	js.Rewrite("$<.GetRecentPaintRequests($1)", last)
	return i
}

// RemoveEventListener
func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	js.Rewrite("$<.RemoveEventListener($1, $2)", eventType, callback)
}

// RepositionWindow
func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	js.Rewrite("$<.RepositionWindow($1, $2)", x, y)
}

// ResizeWindow
func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	js.Rewrite("$<.ResizeWindow($1, $2)", width, height)
}

// ActiveNetworkRequestCount
func (*PerfWidgetExternal) ActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	js.Rewrite("$<.ActiveNetworkRequestCount")
	return activeNetworkRequestCount
}

// AverageFrameTime
func (*PerfWidgetExternal) AverageFrameTime() (averageFrameTime float32) {
	js.Rewrite("$<.AverageFrameTime")
	return averageFrameTime
}

// AveragePaintTime
func (*PerfWidgetExternal) AveragePaintTime() (averagePaintTime float32) {
	js.Rewrite("$<.AveragePaintTime")
	return averagePaintTime
}

// ExtraInformationEnabled
func (*PerfWidgetExternal) ExtraInformationEnabled() (extraInformationEnabled bool) {
	js.Rewrite("$<.ExtraInformationEnabled")
	return extraInformationEnabled
}

// IndependentRenderingEnabled
func (*PerfWidgetExternal) IndependentRenderingEnabled() (independentRenderingEnabled bool) {
	js.Rewrite("$<.IndependentRenderingEnabled")
	return independentRenderingEnabled
}

// IrDisablingContentString
func (*PerfWidgetExternal) IrDisablingContentString() (irDisablingContentString string) {
	js.Rewrite("$<.IrDisablingContentString")
	return irDisablingContentString
}

// IrStatusAvailable
func (*PerfWidgetExternal) IrStatusAvailable() (irStatusAvailable bool) {
	js.Rewrite("$<.IrStatusAvailable")
	return irStatusAvailable
}

// MaxCPUSpeed
func (*PerfWidgetExternal) MaxCPUSpeed() (maxCpuSpeed uint) {
	js.Rewrite("$<.MaxCPUSpeed")
	return maxCpuSpeed
}

// PaintRequestsPerSecond
func (*PerfWidgetExternal) PaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	js.Rewrite("$<.PaintRequestsPerSecond")
	return paintRequestsPerSecond
}

// PerformanceCounter
func (*PerfWidgetExternal) PerformanceCounter() (performanceCounter uint64) {
	js.Rewrite("$<.PerformanceCounter")
	return performanceCounter
}

// PerformanceCounterFrequency
func (*PerfWidgetExternal) PerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	js.Rewrite("$<.PerformanceCounterFrequency")
	return performanceCounterFrequency
}
