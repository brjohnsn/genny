// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package syntax

import "time"

type timeSpanList []time.Duration

type TimeSpanUppercase time.Duration

var _ time.Duration
var timeSpanVariable string

func _() {
	var _ []time.Duration
	var _ timeSpanList
	var _ []timeSpanList
}

func PrintTimeSpan(_timeSpan time.Duration) {
	var _ TimeSpanUppercase

	println(timeSpanVariable, _timeSpan, time.Duration(123))
}

type fractionalList []float64

type FractionalUppercase float64

var _ float64
var fractionalVariable string

func _() {
	var _ []float64
	var _ fractionalList
	var _ []fractionalList
}

func PrintFractional(_fractional float64) {
	var _ FractionalUppercase

	println(fractionalVariable, _fractional, float64(123))
}
