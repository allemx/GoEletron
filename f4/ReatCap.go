// Reatância capacitiva
// https://go.dev/play/p/w_f6pAirSHp
// Resultado: 2.65 KΩ
package main

import (
	"fmt"
	"math"
)

func main() {
	c1 := 1e-06 // 1 microfarad
	f1 := 60.0  // 60 hertz
	Xc := ReatCap(f1, c1)
	fmt.Printf("%0.2f KΩ", Xc/1000)
}
func ReatCap(f, c float64) float64 {
	pi := math.Pi
	xc := 1 / (2 * pi * f * c)
	return xc
}
