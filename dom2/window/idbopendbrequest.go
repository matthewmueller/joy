package window

import "github.com/matthewmueller/golly/js"

// js:"IDBOpenDBRequest,omit"
type IDBOpenDBRequest struct {
	IDBRequest
}

// Onblocked
func (*IDBOpenDBRequest) Onblocked() (onblocked func(Event)) {
	js.Rewrite("$<.Onblocked")
	return onblocked
}

// Onblocked
func (*IDBOpenDBRequest) SetOnblocked(onblocked func(Event)) {
	js.Rewrite("$<.Onblocked = $1", onblocked)
}

// Onupgradeneeded
func (*IDBOpenDBRequest) Onupgradeneeded() (onupgradeneeded func(*IDBVersionChangeEvent)) {
	js.Rewrite("$<.Onupgradeneeded")
	return onupgradeneeded
}

// Onupgradeneeded
func (*IDBOpenDBRequest) SetOnupgradeneeded(onupgradeneeded func(*IDBVersionChangeEvent)) {
	js.Rewrite("$<.Onupgradeneeded = $1", onupgradeneeded)
}
