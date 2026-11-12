package main

import (
	"fmt"
	"sync"
)

func afiseaza(ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	mesaj := <-ch
	fmt.Println("Mesaj primit in goroutine:", mesaj)
}

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup
	wg.Add(1)
	go afiseaza(ch, &wg)

	ch <- "salut"

	wg.Wait()
}
