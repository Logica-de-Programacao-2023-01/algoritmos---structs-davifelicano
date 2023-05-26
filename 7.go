// Crie uma struct chamada Animal com os campos "nome", "espécie", "idade" e "som".
// Escreva funções que permitam modificar o som que o animal faz e uma função que imprima as informações do animal e o som que ele faz
package main

import (
	"fmt"
)

type Animal struct {
	nome    string
	espécie string
	som     string
}

func modi(a Animal, novosom string) Animal {
	a.som = novosom
	return a
}
func main() {
	resposta := Animal{
		nome:    "ola",
		espécie: "macaco",
		som:     "AAAAUUUUUU",
	}
	novosom := "haaawww"

	f := modi(resposta, novosom)
	fmt.Println(f)
}
