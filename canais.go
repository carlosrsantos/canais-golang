package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	// go especifica uma Goroutine
	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for {
		//mensagem recebe o valor de canal, obrigatório
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	fmt.Println("Fim da execução")

}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
