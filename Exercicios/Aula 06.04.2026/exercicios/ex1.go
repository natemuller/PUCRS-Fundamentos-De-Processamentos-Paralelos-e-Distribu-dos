// Exercício 1 — Ponto de partida
//
// Este programa usa sync.WaitGroup para a main() esperar
// duas goroutines. Sua tarefa: substituir o WaitGroup por
// canais como mecanismo de sinalização.
package main

import (
	"fmt"
	"time"
)

func crescente(ch chan bool) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("[Crescente] %d\n", i)
		time.Sleep(1 * time.Second)
	}
	ch <- true
}

func decrescente(ch chan bool) {
	for i := 10; i >= 1; i-- {
		fmt.Printf("[Decrescente] %d\n", i)
		time.Sleep(1 * time.Second)
	}
	ch <- true
}

func main() {
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go crescente(ch1)
	go decrescente(ch2)

	<-ch1
	<-ch2
	fmt.Println("Fim!")
}
