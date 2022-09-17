// Função calcula freq ressonancia de um
// circuito LC
// Resultado:5.033069648525193 Mhz
// https://go.dev/play/p/pUgQPVgWVga

package main

import (
	"fmt"
	"math"
)

func main() {
	l1 := 100e-06 // 100 microhenrys
	c1 := 10e-12  // 10 pf
	res := Ressonancia(l1, c1)
	fmt.Println(res/1000000, "Mhz")
}
func Ressonancia(l, c float64) float64 {
	rq := math.Sqrt(l * c)
	f := 1 / (2 * 3.1415 * rq)
	return f
}
