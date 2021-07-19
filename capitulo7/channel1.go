package capitulo7

import (
	"fmt"
	"time"
)

func produzir(c chan int) {
	time.Sleep(2 * time.Second)
	c <- 35
}

func ExecutarChannel1() {
	c := make(chan int)
	go produzir(c)

	valor := <-c
	fmt.Println(valor)
}
