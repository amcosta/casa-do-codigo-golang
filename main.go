package main

import (
	"cdc-golang/capitulo5"
	"cdc-golang/capitulo6"
	"cdc-golang/capitulo7"
	"fmt"
)

func main() {
	// cap5()
	// cap6()
	cap7()
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

	// capitulo6.ServerInit()
}

func cap7() {
	// capitulo7.IniciandoGoroutine()
	// capitulo7.ExecutarChannel1()
	// capitulo7.ExemploBuffer1()
	// capitulo7.ExemploSepararNumerosImparEPar()
	// capitulo7.ExemploComTimeout()
	// capitulo7.ExemploWaitGroup()
	capitulo7.GoMaxProc()
}
