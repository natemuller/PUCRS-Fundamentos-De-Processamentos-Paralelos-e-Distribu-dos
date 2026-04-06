package main

import (
	"fmt"
	"time"
)

func produtor(ch chan<- string) {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("mensagem %d", i)
		fmt.Printf("[Produtor]   enviando %s\n", msg)
		ch <- msg
		time.Sleep(200 * time.Millisecond)
	}
	close(ch)
	fmt.Println("[Produtor] Canal fechado")
}

func consumidor(ch <-chan string, done chan<- bool) {
	for msg := range ch {
		horario := time.Now().Format("15:04:05")
		fmt.Printf("[Consumidor] Recebeu %s às %s\n", msg, horario)
		time.Sleep(1 * time.Second)
	}
	done <- true
}

func main() {
	ch := make(chan string, 2)
	done := make(chan bool)

	go produtor(ch)
	go consumidor(ch, done)

	<-done

	/*
		a) ch := make(chan string)
		b) Porque no canal sem buffer, o produtor só consegue enviar quando o consumidor estiver pronto para receber.
		   Então o ch <- msg pode ficar bloqueado esperando.
		c) ch := make(chan string, 5)
		d) Com buffer 5, o produtor consegue colocar as 5 mensagens no canal sem precisar esperar a consumidora receber
		   na mesma hora.
		e) O produto bloqueia.
	*/
}
