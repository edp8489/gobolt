package utils

import (
	"fmt"
	"math"
	"testing"
)

func TestCalcPreloadImperial(t *testing.T) {
	/*
		Set up unit test for  imperial units
		D = 0.250 in
		nominal torque = 80 in-lbs
		Torque tolerance = +/- 5 in-lbs
		Preload uncertainty = +/- 35%
		Nut factor K = 0.2
	*/
	var want preload
	want.Nominal = 1600.0
	want.Max = 2295.0
	want.Min = 975.0
	want.Units = "imperial"

	fmt.Println("Testing preload function, imperial units...")

	got := CalcPreload(80.0, 0.25, 5.0, 0.2, 0.35, "in-lbf")

	if math.Abs(got.Nominal-want.Nominal) > 1 || math.Abs(got.Min-want.Min) > 1 || math.Abs(got.Max-want.Max) > 1 {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCalcPreloadMetric(t *testing.T) {
	/*
		Set up unit test for  imperial units
		D = 6 mm
		nominal torque = 10 Nm
		Torque tolerance = +/- 0.5 Nm
		Preload uncertainty = +/- 35%
		Nut factor K = 0.2
	*/
	var want preload
	want.Nominal = 8333.33
	want.Max = 11812.5
	want.Min = 5145.8
	want.Units = "metric"

	fmt.Println("Testing preload function, metric units...")

	got := CalcPreload(10, 6.0, 0.5, 0.2, 0.35, "Nm")

	if math.Abs(got.Nominal-want.Nominal) > 1 || math.Abs(got.Min-want.Min) > 1 || math.Abs(got.Max-want.Max) > 1 {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
