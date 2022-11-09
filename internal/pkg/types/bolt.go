package types

import (
	"math"
)

type Bolt struct {
	DIA_MAJ      float64
	Pitch        int
	Thread_class int8
	Series       string
	Material     Material
	Fsu          float64
	Grip         float64
	Dminor       float64
	Dpitch       float64
	Amaj         float64
	Amin         float64
	Athrd        float64
}

// @TODO refactor Bolt properties into a Thread type
type Thread struct {
	DIA_MIN  float64
	DIA_PTCH float64
	A_TEN    float64
	A_maj    float64
	A_min    float64
	A_thrd   float64
}

// @todo constructor factory methods
// ref https://go.dev/doc/effective_go#allocation_new
func NewBolt(d float64) *Bolt {
	b := new(Bolt)
	b.DIA_MAJ = d
	b.Amaj = Area(b.DIA_MAJ)
	return b
}

// @todo, return error if improper thread series string is input
func MinorDiam_UN(d float64, n int64) float64 {
	p := float64(1 / n)
	dmin := d - 1.299038*p
	return dmin
}

func MinorDiam_M(d float64, p float64) float64 {
	dmin := 1.226869 * p
	return dmin
}

// @TODO return error for invalid params
// @TODO verify if eq applies to both UN and M thread series
func PitchDia(d float64, p float64) float64 {
	dp := d - 0.649519*p
	return dp
}

func ThreadShearAreaExt(b *Bolt, len float64, int_ext string) float64 {
	// @todo
	// area is based on thread engagement (= nut or insert length)
	// Internal: FED-STD-H28/2B, Table II.B.1, eq 2
	// External: FED-STD-H28/2B, Table II.B.1, eq 4
	return 0.0
}

func TensileStressArea(dpitch float64, dminor float64) float64 {
	// Returns the tensile stress area, which is calculated midway between
	// the minor diameter and pitch diameter
	davg := (dpitch + dminor) / 2
	return Area(davg)
}

func Area(d float64) float64 {
	A := math.Pi * math.Pow(d, 2) / 4
	return A
}
