package types

import (
	"errors"
)

type BearingEntry struct {
	Basis    string
	edgeDist float64
	Fbru     float64
	Fbry     float64
}

type CompressionEntry struct {
	Basis    string
	grainDir string
	Fcu      float64
	Fcy      float64
}

type ShearEntry struct {
	Basis    string
	grainDir string
	Fsu      float64
}

type TensionEntry struct {
	Basis    string
	grainDir string
	Ftu      float64
	Fty      float64
}

type Material struct {
	Name            string
	UnitSystem      string
	Source          string
	Modulus         float64
	Poisson         float64
	BearingData     []BearingEntry
	CompressionData []CompressionEntry
	ShearData       []ShearEntry
	TensionData     []TensionEntry
}

func GetAllowable(mat *Material, prop string, basis string, dir string) (float64, error) {
	// return error if user is requesting bearing allowable
	if prop == "Fbru" || prop == "Fbry" {
		return 0, errors.New("for bearing data use function (tbd)")
	}

	var value float64
	var err error = nil
	switch prop {
	case "Ftu":
		// todo
	case "Fty":
		//todo
	case "Fcy":
		//todo
	case "Fsu":
		//todo
	default:
		value = 0
		err = errors.New("invalid property requested. Options are 'Ftu', 'Fty', 'Fcy', 'Fsu'")

	}
	return value, err
}
