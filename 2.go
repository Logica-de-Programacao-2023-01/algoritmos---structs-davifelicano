// Crie uma struct chamada Pessoa com os campos "nome", "idade" e "endereço".
// O campo "endereço" deve ser uma outra struct com os campos "rua", "número", "cidade" e "estado".
// Escreva uma função que receba uma Pessoa como parâmetro e imprima seu endereço completo
package main

import "fmt"

type endereço struct {
	rua    string
	número int
	cidade string
	estado string
}
type pessoa struct {
	nome     string
	idade    int
	endereço endereço
}

func rec(p pessoa) {
	fmt.Println("Rua", p.endereço.rua)
	fmt.Println("Número", p.endereço.número)
	fmt.Println("cidade", p.endereço.cidade)
	fmt.Println("Estado", p.endereço.estado)
}

func main() {
	pe := pessoa{
		nome:  "Davi",
		idade: 18,
		endereço: endereço{
			rua:    "guara",
			número: 111,
			cidade: "Brasília",
			estado: "DF",
		},
	}
	rec(pe)

}
