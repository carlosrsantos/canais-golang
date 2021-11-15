package main

import "fmt"

//worker pools é um padrão de concorrência para a execução de várias tarefas simultânea

func main() {
	//canal<- = canal que envia um valor
	tarefas := make(chan int, 45)

	//<-canal = canal que recebe um valor
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

//canal<- = canal que envia um valor
//<-canal = canal que recebe um valor
//mais informações: http://goporexemplo.golangbr.org/channel-directions.html
func worker(tarefas <-chan int, resultados chan<- int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
