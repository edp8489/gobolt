package types

import "testing"

// test minor diameter (UN thread)
func TestMinorDiamUN(t *testing.T) {
	// 0.250-28 UNF 2A

	d := 0.250 // in
	n = 28     // unitless

	got := MinorDiam_UN(d, n)

	//want :=
}

// test minor diameter (M thread)

// test minor diameter (UNJ thread)

// test pitch diameter (UN thread)
// 0.250-28 UNF 2A
// dpitch = 0.2268 per ASME B1.1 Table 7

// test pitch diameter (M thread)

// test pitch diameter (UNJ thread)

// test tensile stress area (UN thread)
// 0.250-28 UNF 2A
// A = 0.0364 in**2 per ASME B1.1 Table 7

// test tensile stress area (M thread)

// test tensile stress area (UNJ thread)
