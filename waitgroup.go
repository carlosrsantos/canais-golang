package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Goroutine 1")
		waitGroup.Done() // -1 .Add
	}()

	go func() {
		escrever("Goroutine 2")
		waitGroup.Done() // -1 .Add
	}()

	go func() {
		escrever("Goroutine 3")
		waitGroup.Done() // -1 .Add
	}()

	go func() {
		escrever("Goroutine 4")
		waitGroup.Done() // -1 .Add
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
