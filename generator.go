package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i <= 10; i++ {
		canal := escrever(i)
		fmt.Println(<-canal)
	}
}

func escrever(numero int) <-chan int {
	canal := make(chan int)

	go func() {
		for {
			// canal <- fmt.Sprintf("Valor recebido: %s", texto)
			canal <- numero
			time.Sleep(time.Second * 1)
		}
	}()

	return canal
}
