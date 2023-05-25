// Crie uma struct chamada Circulo com o campo "raio".
// Escreva uma função que receba um Circulo como parâmetro e calcule a área do círculo (área = pi * raio * raio).
// Dica: use a constante math.Pi para representar o número pi.
package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func cir(c circulo) float64 {
	area := math.Pi * c.raio * c.raio
	return area
}
func main() {
	ar := circulo{
		raio: 10,
	}
	fmt.Print(cir(ar))

}
