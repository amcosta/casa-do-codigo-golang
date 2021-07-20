package capitulo7

import (
	"fmt"
	"time"
)

func timeout(c chan<- bool) {
	time.Sleep(5 * time.Second)
	c <- true
}

func ExemploComTimeout() {
	c := make(chan bool, 1)
	go timeout(c)

	fmt.Println("Executando...")

	fim := false
	for !fim {
		select {
		case fim = <-c:
			fmt.Println("Fim!")
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout!")
			fim = true
		}
	}
}
