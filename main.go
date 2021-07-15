package main

import (
	"cdc-golang/capitulo5"
	"cdc-golang/capitulo6"
	"fmt"
)

func main() {
	cap5()
	cap6()
}

func cap5() {
	fmt.Println("\n\nExercícios do Capitulo 5")
	capitulo5.Executar()
}

func cap6() {
	fmt.Println("\n\nExercícios do Capitulo 6")
	precoDolar, precoReal := capitulo6.PrecoFinal(34.99)
	fmt.Printf("Preço final em dólar: %.2f\n", precoDolar)
	fmt.Printf("Preço final em reais: %.2f\n", precoReal)

	capitulo6.CriarArquivos()
	capitulo6.FuncaoAnonima()
	capitulo6.Fibonacci()
	capitulo6.Cronometrar(capitulo6.GerarFibonacci(8))
	capitulo6.Cronometrar(capitulo6.GerarFibonacci(48))
	capitulo6.Cronometrar(capitulo6.GerarFibonacci(88))
}
