// Calc impedância RLC série 
// Resultado: 47.0 ohms 
// https://go.dev/play/p/j6Z7nMF20VV

package main import ( "fmt"
                      "math" ) 
func main() {
	r1 := 47.0 rl1 := 50.0 
       rc1 := 50.0 Z := ZRLCserie(r1, rl1, rc1) fmt.Printf("%0.1f ohms", Z)
}
func ZRLCserie(r, rl, rc float64) float64 {
	z := math.Sqrt(r*r + (rl*rl - rc*rc))
             return z
}
