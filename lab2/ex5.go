package main

import (
	"fmt"
	"sync"
)

func send(ch chan<- string) {
	ch <- "salut"
}

func receive(ch <-chan string) string {
	return <-ch
}

func main() {
	ch := make(chan string)
	var mesaj string

	var wg sync.WaitGroup
	wg.Add(1)

	go send(ch)
	go func() {
		defer wg.Done()
		mesaj = receive(ch)
	}()

	wg.Wait()
	fmt.Println("Mesaj primit:", mesaj)
}
