package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	canal := multiplex(escrever("Hoje chove"), escrever("Amanhã faz calor"))

	for i := 0; i <= 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplex(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case msg := <-canalEntrada1:
				canalSaida <- msg
			case msg := <-canalEntrada2:
				canalSaida <- msg
			}
		}
	}()

	return canalSaida
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
