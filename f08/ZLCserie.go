// Calc impedância LC serie
// Resultado: 0 ohms 
// https://go.dev/play/p/Pq_ZC1DBIVR

package main 

import "fmt"

 func main() { 
	xl1 := 50.0 // Res indutiva 
        xc1 := 50.0 // Res capacitiva
	Z1 := ZLCserie(xl1, xc1) 
	fmt.Printf("%f ohms", Z1)
}
func ZLCserie(xc, xl float64) float64 { 
	z := xl - xc
             return z
}
