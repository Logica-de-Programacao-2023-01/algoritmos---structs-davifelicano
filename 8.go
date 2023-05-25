// Crie uma struct chamada Viagem com os campos "origem", "destino", "data" e "preço".
// Escreva uma função que receba um slice de Viagens como parâmetro e retorne a viagem mais cara.
package main

import "fmt"

type Viagem struct {
	origem  string
	destino string
	data    string
	preço   float64
}

func maior(v []Viagem) Viagem {
	maiscara := v[0]
	for _, vv := range v {
		if vv.preço > maiscara.preço {
			maiscara = vv
		}
	}
	return maiscara
}
func main() {
	resultado := []Viagem{
		{origem: "maringa", destino: "DF", data: "2/2/2004", preço: 1000},
		{origem: "Mg", destino: "RJ", data: "25/05/2023", preço: 1500},
	}
	f := maior(resultado)
	fmt.Print(f)
}
