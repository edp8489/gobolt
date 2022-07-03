package types

// @todo: logic to handle default values (e.g. Factors, where default shuld be 1 rather than 0)

type Factors struct {
	SFult float64
	SFyld float64
	FF    float64
}

type Loadset struct {
	Ptension float64
	Pshear   float64
	Factors  Factors
}

type Bolt struct {
	D            float64
	Pitch        int
	Thread_class int8
	Material     Material
	Fsu          float64
}

type Member struct {
	T        float64
	Edge     float64
	EoverD   float64
	Fbru     float64
	Fbry     float64
	Material Material
}
