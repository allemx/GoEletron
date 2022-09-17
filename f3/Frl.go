// Função calc Freq RL Resultado: 
// 1.5915494309189533 Mhz 
// 1591.5494309189532 Khz

package main

 import ( "fmt"
          "math"
        ) 
func main() {
	r := 1e+03 // 1 K ohms
        l :=10e-06 // 100 microhenry
        fm := FreqRL(r, l) / 10e+06
        fk := FreqRL(r, l) / 10e+03 
	
        fmt.Println(fm, "Mhz") 
	fmt.Println(fk, "Khz")
}
func FreqRL(r, l float64) float64 { 
        pi := math.Pi
	 f := r / (2 * pi * l)
           return f
}
