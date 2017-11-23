package window

import "github.com/matthewmueller/golly/js"

// IDBOpenDBRequest struct
// js:"IDBOpenDBRequest,omit"
type IDBOpenDBRequest struct {
	IDBRequest
}

// Onblocked prop
func (*IDBOpenDBRequest) Onblocked() (onblocked func(Event)) {
	js.Rewrite("$<.onblocked")
	return onblocked
}

// Onblocked prop
func (*IDBOpenDBRequest) SetOnblocked(onblocked func(Event)) {
	js.Rewrite("$<.onblocked = $1", onblocked)
}

// Onupgradeneeded prop
func (*IDBOpenDBRequest) Onupgradeneeded() (onupgradeneeded func(*IDBVersionChangeEvent)) {
	js.Rewrite("$<.onupgradeneeded")
	return onupgradeneeded
}

// Onupgradeneeded prop
func (*IDBOpenDBRequest) SetOnupgradeneeded(onupgradeneeded func(*IDBVersionChangeEvent)) {
	js.Rewrite("$<.onupgradeneeded = $1", onupgradeneeded)
}
