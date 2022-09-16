// Exemplo de calculo filtro passivo RC
// Frequencia de corte
// Resistor em Ohms e Capacitor em farads
// Resultado: 60 hertz

package main

import "fmt"

func main() {
	r1 := 2654.0 // 2.654 K ohms
	c1 := 1e-06  // 1 microfarad
	fc := Frequencia(r1, c1)
	fmt.Printf("%0.1f hertz", fc)
}
func Frequencia(r, c float64) float64 {
	f := 1 / (2 * 3.14 * r * c)
	return f
}
