package main

import (
	"fmt"
)

func main() {

	//Canais com bufffer
	canal := make(chan string, 3)
	//enviar e receber valores são funções bloqueantes
	canal <- "Canal 1"
	canal <- "Canal 2"
	canal <- "Canal 3"

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3)

}
