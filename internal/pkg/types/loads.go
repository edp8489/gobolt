package types

// @TODO: constructor factory method w/ logic to handle default values
// e.g. all factors should default to 1, not 0

type Factors struct {
	SFult float64
	SFyld float64
	FF    float64
	SFsep float64
}

type Loadset struct {
	Ptension float64
	Pshear   float64
	Factors  Factors
}
