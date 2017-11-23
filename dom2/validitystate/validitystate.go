package validitystate

import "github.com/matthewmueller/golly/js"

// ValidityState struct
// js:"ValidityState,omit"
type ValidityState struct {
}

// BadInput
func (*ValidityState) BadInput() (badInput bool) {
	js.Rewrite("$<.BadInput")
	return badInput
}

// CustomError
func (*ValidityState) CustomError() (customError bool) {
	js.Rewrite("$<.CustomError")
	return customError
}

// PatternMismatch
func (*ValidityState) PatternMismatch() (patternMismatch bool) {
	js.Rewrite("$<.PatternMismatch")
	return patternMismatch
}

// RangeOverflow
func (*ValidityState) RangeOverflow() (rangeOverflow bool) {
	js.Rewrite("$<.RangeOverflow")
	return rangeOverflow
}

// RangeUnderflow
func (*ValidityState) RangeUnderflow() (rangeUnderflow bool) {
	js.Rewrite("$<.RangeUnderflow")
	return rangeUnderflow
}

// StepMismatch
func (*ValidityState) StepMismatch() (stepMismatch bool) {
	js.Rewrite("$<.StepMismatch")
	return stepMismatch
}

// TooLong
func (*ValidityState) TooLong() (tooLong bool) {
	js.Rewrite("$<.TooLong")
	return tooLong
}

// TypeMismatch
func (*ValidityState) TypeMismatch() (typeMismatch bool) {
	js.Rewrite("$<.TypeMismatch")
	return typeMismatch
}

// Valid
func (*ValidityState) Valid() (valid bool) {
	js.Rewrite("$<.Valid")
	return valid
}

// ValueMissing
func (*ValidityState) ValueMissing() (valueMissing bool) {
	js.Rewrite("$<.ValueMissing")
	return valueMissing
}
