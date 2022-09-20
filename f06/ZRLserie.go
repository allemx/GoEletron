// Impedância circuito RL série 
// Resultado: 68.62 ohms 
// https://go.dev/play/p/__QqHOtfibW
package main

 import ( "fmt"     
          "math" ) 
func main() {
	xl1 := 50.0 // reatância indutiva
         r1 := 47.0 // resistência serie
         z1 := ZRLserie(xl1, r1) 
	fmt.Printf("%0.2f ohms", z1)
}
func ZRLserie(xl, r float64) float64 {
          z := math.Sqrt(xl*xl + r*r)
	    return z
}
