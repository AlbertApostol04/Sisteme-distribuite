package main

import (
	"fmt"
	"time"
)

func lucreaza(done chan<- bool) {
	fmt.Println("Goroutine: incep procesarea")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Goroutine: procesarea s-a terminat")

	done <- true
}

func main() {
	done := make(chan bool)

	go lucreaza(done)

	<-done
	fmt.Println("main: am primit notificarea, goroutine-ul a terminat")
}
