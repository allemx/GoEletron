// Impedância circuito RC série 
// Resultado: 68.62 ohms 
// https://go.dev/play/p/ApxAxhtmj9Q

package main import ( "fmt"
                     "math" ) 
func main() {
	xc1 := 50.0 // reatância capacitiva 
         r1 := 47.0 // resistência serie 
         z1 := ZRCserie(xc1, r1) 
	       fmt.Printf("%0.2f ohms", z1)
}
func ZRCserie(xc, r float64) float64 { 
z := math.Sqrt(xc*xc + r*r)
	return z
}
