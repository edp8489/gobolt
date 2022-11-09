package output

import (
	"fmt"

	"github.com/edp8489/gobolt/internal/pkg/types"
)

func PrintLoads(l types.Loadset) {
	fmt.Println("===Loadset===")
	fmt.Println("Tension load: ", l.Ptension)
	fmt.Println("Shear load: ", l.Pshear)
	fmt.Println("Safety Factors: ")
	fmt.Println("  SFult: ", l.Factors.SFult)
	fmt.Println("  SFyld: ", l.Factors.SFyld)
	fmt.Println("  Fitting: ", l.Factors.FF)

}

func PrintBolt(b types.Bolt) {
	fmt.Println("===Bolt Properties===")
	fmt.Println("Diameter: ", b.DIA_MAJ, " [in]")
	fmt.Println("Material: ", b.Fsu, " [psi] Fsu")
}
