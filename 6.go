// Crie uma struct chamada Funcionário com os campos "nome", "salário" e "idade".
// Escreva funções que permitam aumentar ou diminuir o salário do funcionário em uma determinada porcentagem
// e uma função que calcule o tempo de serviço do funcionário (considerando que ele começou a trabalhar aos 18 anos).
package main

import "fmt"

type Funcionário struct {
	nome    string
	salário float64
	idade   int
}

func aumentar(w Funcionário) float64 {
	var valor float64
	valor = 1.1
	ad := w.salário * valor
	return ad
}
func tempo(s Funcionário) int {
	te := s.idade - 18
	return te
}
func main() {
	resultado := Funcionário{
		nome:    "Davi",
		salário: 10.000,
		idade:   28,
	}

	novosalario := aumentar(resultado)
	fmt.Println(novosalario)
	tempodeserviço := tempo(resultado)
	fmt.Println(tempodeserviço)
}
