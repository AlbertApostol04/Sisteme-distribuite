package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Sisteme"
		ch <- "Distribuite"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
