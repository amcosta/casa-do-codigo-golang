package capitulo7

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func ExemploWaitGroup() {
	fmt.Println(runtime.NumCPU())

	inicio := time.Now()
	rand.Seed(inicio.UnixNano())

	var controle sync.WaitGroup

	for i := 0; i < 5; i++ {
		controle.Add(1)
		go wgExecutar(&controle)
	}

	controle.Wait()
	fmt.Printf("Finalizado em %s.\n", time.Since(inicio))
}

func wgExecutar(controle *sync.WaitGroup) {
	defer controle.Done()
	duracao := time.Duration(1+rand.Intn(5)) * time.Second
	fmt.Printf("Dormindo por %s...\n", duracao)
	time.Sleep(duracao)
}
