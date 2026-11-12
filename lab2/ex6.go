package main

import (
	"fmt"
	"math/rand"
	"time"
)

func robot(id int, finish chan<- int) {
	for pas := 0; pas < 10; pas++ {
		time.Sleep(time.Duration(10+rand.Intn(40)) * time.Millisecond)
	}
	finish <- id
}

func main() {
	finish1 := make(chan int, 1)
	finish2 := make(chan int, 1)

	go robot(1, finish1)
	go robot(2, finish2)

	select {
	case id := <-finish1:
		fmt.Println("Robotul", id, "a ajuns primul la finish")
	case id := <-finish2:
		fmt.Println("Robotul", id, "a ajuns primul la finish")
	}
}
