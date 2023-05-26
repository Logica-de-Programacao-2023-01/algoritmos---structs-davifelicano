package main

import "fmt"

type retangulo struct {
	altura  float64
	largura float64
}

func area(p retangulo) float64 {
	calculo := p.altura * p.largura
	return calculo
}
func main() {
	var altura, largura float64
	fmt.Print("Informe a altura")
	fmt.Scan(&altura)
	fmt.Print("Informe a altura")
	fmt.Scan(&largura)
	dados := retangulo{
		largura: largura,
		altura:  altura,
	}
	calcul := area(dados)
	fmt.Print(calcul)
}
