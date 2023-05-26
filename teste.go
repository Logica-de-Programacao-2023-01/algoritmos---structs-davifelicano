package main

import "fmt"

type Pessoa struct {
	nome   string
	idade  int
	altura float64
}

func criaPessoa(nome string, idade int, altura float64) Pessoa {

	return Pessoa{nome: nome, idade: idade, altura: altura}

}
func main() {
	p := criaPessoa("Maria", 25, 1.65)
	c := Pessoa{nome: "João", idade: 30, altura: 1.75}
	fmt.Println(p)
	fmt.Print(c.nome)
	type Endereco struct {
		rua    string
		cidade string
		estado string
	}

	type Pessoa struct {
		nome   string
		idade  int
		altura float64
	}
	p := Pessoa{

		nome: "João"
		,

		idade: 30,
		endereco: Endereco{
		rua: "Rua X"
		,

		cidade: "São Paulo"
		,

		estado: "SP"
		,

	},

	}
}
