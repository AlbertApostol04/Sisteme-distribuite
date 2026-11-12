package main

import (
	"fmt"
	"time"
)

type Pos struct{ X, Y int }

func robot(id int, traseu []Pos, out chan<- Pos, in <-chan Pos, raport chan<- string) {
	miscare := make(chan bool)

	go func() {
		for _, p := range traseu {
			out <- p
			fmt.Printf("Robotul %d se afla la (%d, %d)\n", id, p.X, p.Y)
			time.Sleep(100 * time.Millisecond)
		}
		close(out)
		miscare <- true
	}()

	for p := range in {
		fmt.Printf("Robotul %d a primit pozitia (%d, %d) a celuilalt robot\n", id, p.X, p.Y)
	}

	<-miscare
	raport <- fmt.Sprintf("Robotul %d si-a terminat sarcina", id)
}

func main() {
	c12 := make(chan Pos)
	c21 := make(chan Pos)
	raport := make(chan string)

	traseu1 := []Pos{{0, 0}, {1, 0}, {2, 1}, {3, 2}, {4, 3}}
	traseu2 := []Pos{{9, 9}, {8, 8}, {7, 7}}

	go robot(1, traseu1, c12, c21, raport)
	go robot(2, traseu2, c21, c12, raport)

	for i := 0; i < 2; i++ {
		fmt.Println(<-raport)
	}
	fmt.Println("Observer: ambii roboti si-au terminat sarcina")
}
