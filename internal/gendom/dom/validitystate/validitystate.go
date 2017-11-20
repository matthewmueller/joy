package validitystate

import "github.com/matthewmueller/golly/js"

type ValidityState struct {
}

func (*ValidityState) GetBadInput() (badInput bool) {
	js.Rewrite("$<.badInput")
	return badInput
}

func (*ValidityState) GetCustomError() (customError bool) {
	js.Rewrite("$<.customError")
	return customError
}

func (*ValidityState) GetPatternMismatch() (patternMismatch bool) {
	js.Rewrite("$<.patternMismatch")
	return patternMismatch
}

func (*ValidityState) GetRangeOverflow() (rangeOverflow bool) {
	js.Rewrite("$<.rangeOverflow")
	return rangeOverflow
}

func (*ValidityState) GetRangeUnderflow() (rangeUnderflow bool) {
	js.Rewrite("$<.rangeUnderflow")
	return rangeUnderflow
}

func (*ValidityState) GetStepMismatch() (stepMismatch bool) {
	js.Rewrite("$<.stepMismatch")
	return stepMismatch
}

func (*ValidityState) GetTooLong() (tooLong bool) {
	js.Rewrite("$<.tooLong")
	return tooLong
}

func (*ValidityState) GetTypeMismatch() (typeMismatch bool) {
	js.Rewrite("$<.typeMismatch")
	return typeMismatch
}

func (*ValidityState) GetValid() (valid bool) {
	js.Rewrite("$<.valid")
	return valid
}

func (*ValidityState) GetValueMissing() (valueMissing bool) {
	js.Rewrite("$<.valueMissing")
	return valueMissing
}
