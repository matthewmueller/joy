package validitystate

import "github.com/matthewmueller/joy/macro"

// ValidityState struct
// js:"ValidityState,omit"
type ValidityState struct {
}

// BadInput prop
// js:"badInput"
func (*ValidityState) BadInput() (badInput bool) {
	macro.Rewrite("$_.badInput")
	return badInput
}

// CustomError prop
// js:"customError"
func (*ValidityState) CustomError() (customError bool) {
	macro.Rewrite("$_.customError")
	return customError
}

// PatternMismatch prop
// js:"patternMismatch"
func (*ValidityState) PatternMismatch() (patternMismatch bool) {
	macro.Rewrite("$_.patternMismatch")
	return patternMismatch
}

// RangeOverflow prop
// js:"rangeOverflow"
func (*ValidityState) RangeOverflow() (rangeOverflow bool) {
	macro.Rewrite("$_.rangeOverflow")
	return rangeOverflow
}

// RangeUnderflow prop
// js:"rangeUnderflow"
func (*ValidityState) RangeUnderflow() (rangeUnderflow bool) {
	macro.Rewrite("$_.rangeUnderflow")
	return rangeUnderflow
}

// StepMismatch prop
// js:"stepMismatch"
func (*ValidityState) StepMismatch() (stepMismatch bool) {
	macro.Rewrite("$_.stepMismatch")
	return stepMismatch
}

// TooLong prop
// js:"tooLong"
func (*ValidityState) TooLong() (tooLong bool) {
	macro.Rewrite("$_.tooLong")
	return tooLong
}

// TypeMismatch prop
// js:"typeMismatch"
func (*ValidityState) TypeMismatch() (typeMismatch bool) {
	macro.Rewrite("$_.typeMismatch")
	return typeMismatch
}

// Valid prop
// js:"valid"
func (*ValidityState) Valid() (valid bool) {
	macro.Rewrite("$_.valid")
	return valid
}

// ValueMissing prop
// js:"valueMissing"
func (*ValidityState) ValueMissing() (valueMissing bool) {
	macro.Rewrite("$_.valueMissing")
	return valueMissing
}
