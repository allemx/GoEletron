// Reatância Indutiva
// Resultado: 376.99 ohms 
// https://go.dev/play/p/w5L36DMip9x

package main 

import ( "fmt"
        "math" ) 

func main() {
	fr1 := 60.0 // 60 hertz 
         l1 :=	1.0 // 1 henry
         Xl :=ReatIndut(fr1, l1) 
	 fmt.Printf("%.02f Ohms", Xl)
}
func ReatIndut(f, l float64) float64 { 
	pi := math.Pi
        xl := 2 * pi * f * l
         return xl
}
