//Crie uma struct chamada Aluno com os campos "nome"
//,
//"idade" e "notas". O campo
//"notas" deve ser um slice de float64. Escreva uma função que receba um Aluno como
//parâmetro e calcule a média das suas notas.
package main

import "fmt"

type aluno struct {
	idade int
	notas []float64
}

func med(a aluno) float64 {
	qn := len(a.notas)
	var soma float64
	for _, x := range a.notas {
		soma += x
	}
	media := soma / float64(qn)
	return media
}
func main() {
	boletim := aluno{
		idade: 18,
		notas: []float64{1, 2, 3, 4, 5, 6, 7},
	}
	resultado := med(boletim)
	fmt.Print(resultado)
}
