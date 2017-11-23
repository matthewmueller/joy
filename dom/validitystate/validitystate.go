package validitystate

import "github.com/matthewmueller/golly/js"

// ValidityState struct
// js:"ValidityState,omit"
type ValidityState struct {
}

// BadInput prop
func (*ValidityState) BadInput() (badInput bool) {
	js.Rewrite("$<.badInput")
	return badInput
}

// CustomError prop
func (*ValidityState) CustomError() (customError bool) {
	js.Rewrite("$<.customError")
	return customError
}

// PatternMismatch prop
func (*ValidityState) PatternMismatch() (patternMismatch bool) {
	js.Rewrite("$<.patternMismatch")
	return patternMismatch
}

// RangeOverflow prop
func (*ValidityState) RangeOverflow() (rangeOverflow bool) {
	js.Rewrite("$<.rangeOverflow")
	return rangeOverflow
}

// RangeUnderflow prop
func (*ValidityState) RangeUnderflow() (rangeUnderflow bool) {
	js.Rewrite("$<.rangeUnderflow")
	return rangeUnderflow
}

// StepMismatch prop
func (*ValidityState) StepMismatch() (stepMismatch bool) {
	js.Rewrite("$<.stepMismatch")
	return stepMismatch
}

// TooLong prop
func (*ValidityState) TooLong() (tooLong bool) {
	js.Rewrite("$<.tooLong")
	return tooLong
}

// TypeMismatch prop
func (*ValidityState) TypeMismatch() (typeMismatch bool) {
	js.Rewrite("$<.typeMismatch")
	return typeMismatch
}

// Valid prop
func (*ValidityState) Valid() (valid bool) {
	js.Rewrite("$<.valid")
	return valid
}

// ValueMissing prop
func (*ValidityState) ValueMissing() (valueMissing bool) {
	js.Rewrite("$<.valueMissing")
	return valueMissing
}
