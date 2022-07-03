package types

import "math"

func ThreadShearAreaExt(b *Bolt) float64 {
	// @todo
	return 0.0
}

func ThreadShearAreaInt(m *Member) float64 {
	// @todo
	return 0.0
}

func Area(d float64) float64 {
	A := math.Pi * math.Pow(d, 2) / 4
	return A
}
