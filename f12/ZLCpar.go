// Func calc impedância LC paralelo 
// Resultado: 150 ohms 
// https://go.dev/play/p/KG04GkiRK0W

package main

import "fmt" 

func main() { 
	xl1 := 75.0 
        xc1 := 50.0 
         Z1 := ZLCpar(xl1, xc1) 
               fmt.Println(Z1,"ohms")
}
func ZLCpar(xl, xc float64) float64 {
      Z := (xl * xc) / (xl - xc)
	return Z
}
