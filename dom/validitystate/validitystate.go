package validitystate

import "github.com/matthewmueller/joy/js"

// ValidityState struct
// js:"ValidityState,omit"
type ValidityState struct {
}

// BadInput prop
// js:"badInput"
func (*ValidityState) BadInput() (badInput bool) {
	js.Rewrite("$_.badInput")
	return badInput
}

// CustomError prop
// js:"customError"
func (*ValidityState) CustomError() (customError bool) {
	js.Rewrite("$_.customError")
	return customError
}

// PatternMismatch prop
// js:"patternMismatch"
func (*ValidityState) PatternMismatch() (patternMismatch bool) {
	js.Rewrite("$_.patternMismatch")
	return patternMismatch
}

// RangeOverflow prop
// js:"rangeOverflow"
func (*ValidityState) RangeOverflow() (rangeOverflow bool) {
	js.Rewrite("$_.rangeOverflow")
	return rangeOverflow
}

// RangeUnderflow prop
// js:"rangeUnderflow"
func (*ValidityState) RangeUnderflow() (rangeUnderflow bool) {
	js.Rewrite("$_.rangeUnderflow")
	return rangeUnderflow
}

// StepMismatch prop
// js:"stepMismatch"
func (*ValidityState) StepMismatch() (stepMismatch bool) {
	js.Rewrite("$_.stepMismatch")
	return stepMismatch
}

// TooLong prop
// js:"tooLong"
func (*ValidityState) TooLong() (tooLong bool) {
	js.Rewrite("$_.tooLong")
	return tooLong
}

// TypeMismatch prop
// js:"typeMismatch"
func (*ValidityState) TypeMismatch() (typeMismatch bool) {
	js.Rewrite("$_.typeMismatch")
	return typeMismatch
}

// Valid prop
// js:"valid"
func (*ValidityState) Valid() (valid bool) {
	js.Rewrite("$_.valid")
	return valid
}

// ValueMissing prop
// js:"valueMissing"
func (*ValidityState) ValueMissing() (valueMissing bool) {
	js.Rewrite("$_.valueMissing")
	return valueMissing
}
