//Crie uma struct chamada Triângulo com os campos "base" e "altura".
//Escreva uma função que receba um Triângulo como parâmetro e calcule a área do triângulo (área = base * altura / 2).

package main

import (
	"fmt"
	"sort"
)

type Triângulo struct {
	base   float64
	altura float64
}

func calculomtdoido(t Triângulo) float64 {
	a := t.base * t.altura / 2
	return a
}
func main() {
	t := Triângulo{
		base:   10,
		altura: 20,
	}
	b := calculomtdoido(t)
	fmt.Print(b)

}
short.slice