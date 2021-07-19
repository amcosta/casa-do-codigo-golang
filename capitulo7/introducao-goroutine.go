package capitulo7

import (
	"fmt"
	"time"
)

func imprimir(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", n)
		time.Sleep(200 * time.Millisecond)
	}
}

func IniciandoGoroutine() {
	go imprimir(2)
	imprimir(3)
}
