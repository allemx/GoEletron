// Calc impedância RL paralelo 
// Resultado: 39.826 ohms 
// https://go.dev/play/p/3OwE8YjfFQu

package main 

import ( "fmt"
         "math" ) 

func main() {
	r1 := 47.0 l1 := 75.0 
        Z1 := ZRLpar(r1, l1) 
         fmt.Printf("%.3f ohms", Z1)
}
func ZRLpar(r, l float64) float64 {
 Z := r * l / (math.Sqrt(r*r + l*l)) 
	return Z
}
