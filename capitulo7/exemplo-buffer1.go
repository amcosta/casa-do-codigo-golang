package capitulo7

import "fmt"

func exBuffer1Produzir(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}

func ExemploBuffer1() {
	c := make(chan int, 3)
	go exBuffer1Produzir(c)

	for valor := range c {
		fmt.Println(valor)
	}
}
