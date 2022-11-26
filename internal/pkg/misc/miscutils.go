package misc

import "math"

// PctErrCheck calculates the percent difference between
// a calculated value, calcVal, and a reference value, refVal,
// and checks if the absolute value is below the specified
// limit, errLim (in decimal form).
func PctErrCheck(calcVal float64, refVal float64, errLim float64) bool {
	var check bool = false

	pct_diff := (calcVal - refVal) / refVal

	if math.Abs(pct_diff) < math.Abs(errLim) {
		check = true
	}
	return check
}
