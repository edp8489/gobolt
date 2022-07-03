package types

type BearingEntry struct {
	e_D  float64
	Fbru float64
	Fbry float64
}

type CompressionEntry struct {
	grainDir string
	Fcu      float64
	Fcy      float64
}

type ShearEntry struct {
	grainDir string
	Fsu      float64
}

type TensionEntry struct {
	grainDir string
	Ftu      float64
	Fty      float64
}

type Material struct {
	Basis           string
	Name            string
	Source          string
	Modulus         float64
	Poisson         float64
	BearingData     []BearingEntry
	CompressionData []CompressionEntry
	ShearData       []ShearEntry
	TensionData     []TensionEntry
}
