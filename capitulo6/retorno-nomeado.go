package capitulo6

func PrecoFinal(precoCusto float64) (precoDolar float64, precoReal float64) {
	fatorLucro := 1.33
	taxaConversao := 2.34

	precoDolar = precoCusto * fatorLucro
	precoReal = precoDolar * taxaConversao

	return
}
