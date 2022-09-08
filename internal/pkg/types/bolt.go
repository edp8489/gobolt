package types

import (
	"fmt"
	"math"
)

// @todo constructor factory method

type Bolt struct {
	D            float64
	Pitch        int
	Thread_class int8
	Series       string
	Material     Material
	Fsu          float64
	Grip         float64
	Dminor       float64
	Dpitch       float64
	Anom         float64
	Amin         float64
	Athrd        float64
}

func NewBolt(d float64) Bolt {
	b := Bolt{D: d}
	b.Anom = Area(b.D)
	return b
}

// @todo, return error if improper thread series string is input
func MinorDia(d float64, p float64, s string) float64 {
	switch s {
	case "Unified":
		dmin := d - 1.299038*p
		return dmin
	case "Metric":
		dmin := 1.226869 * p
		return dmin
	default:
		fmt.Println("ERROR: Invalid thread series string entered!")
		fmt.Println("Valid options are 'Metric' or 'Unified'")
		return 0
	}
}

// @todo return error for invalid params
func PitchDia(d float64, p float64) float64 {
	dp := d - 0.649519*p
	return dp
}

func ThreadShearAreaExt(b *Bolt) float64 {
	// @todo
	return 0.0
}

func Area(d float64) float64 {
	A := math.Pi * math.Pow(d, 2) / 4
	return A
}
