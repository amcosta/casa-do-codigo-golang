package capitulo7

import (
	"fmt"
	"sort"
)

func separarImparPar(nums []int, i, p chan<- int, pronto chan<- bool) {
	for _, n := range nums {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	pronto <- true
}

func ExemploSepararNumerosImparEPar() {
	i, p := make(chan int), make(chan int)
	pronto := make(chan bool)
	nums := []int{2, 23, 42, 5, 8, 6, 7, 4, 99, 100}

	go separarImparPar(nums, i, p, pronto)
	var impares, pares []int
	fim := false

	for !fim {
		select {
		case n := <-i:
			impares = append(impares, n)
		case n := <-p:
			pares = append(pares, n)
		case fim = <-pronto:
		}
	}

	sort.Ints(pares)
	sort.Ints(impares)
	fmt.Printf("Pares: %v\nImpares: %v\n", pares, impares)
}
