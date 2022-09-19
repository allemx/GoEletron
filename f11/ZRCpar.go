// Calc impedância RC paralelo 
// Resultado: 39.826 
// https://go.dev/play/p/Yi65vxcgE1-

package main

import ( "fmt"
          "math" ) 

func main() {
	r1 := 47.0 
        c1 := 75.0 //reat capac 
        Z1 := ZRCpar(r1, c1) 
	   fmt.Printf("%.3f ohms", Z1)
}
func ZRCpar(r, c float64) float64 { 
        Z := r * c / (math.Sqrt(r*r + c*c)) 
	   return Z
}
