// Func calc impedância RLC paralelo 
// Resultado: 44.85 ohms 
// https://go.dev/play/p/RzHPoNCzGAe

package main 

import ( "fmt"
         "math" ) 

func main() {
	r1 := 47.0 
       xl1 := 75.0 
       xc1 := 50.0 
        Z1 := ZRLCpar(r1, xl1, xc1) 
	      fmt.Printf("%.2f ohms", Z1)
}
func ZRLCpar(r, xl, xc float64) float64 {
	Z := (r * xl * xc) / math.Sqrt((xl*xl*xc*xc)+r*r*(xl-xc)*(xl-xc)) 
	      return Z
}
