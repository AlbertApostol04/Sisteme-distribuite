package main

import (
	"fmt"
	"sync"
)

type Msg struct {
	Value int
	Done  bool
}

func node(id int, in <-chan Msg, out chan<- Msg, isLeader *bool, wg *sync.WaitGroup) {
	defer wg.Done()

	out <- Msg{Value: id}

	for {
		m := <-in

		if m.Done {
			if m.Value != id {
				out <- m
			}
			return
		}

		if m.Value > id {
			out <- m
		} else if m.Value == id {
			*isLeader = true
			out <- Msg{Value: id, Done: true}
		}
	}
}

func RingFlooding(ids []int) int {
	n := len(ids)
	if n == 0 {
		return -1
	}

	chans := make([]chan Msg, n)
	for i := range chans {
		chans[i] = make(chan Msg, n+1)
	}

	leaders := make([]bool, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go node(ids[i], chans[i], chans[(i+1)%n], &leaders[i], &wg)
	}
	wg.Wait()

	for i := 0; i < n; i++ {
		if leaders[i] {
			return ids[i]
		}
	}
	return -1
}

func main() {
	ids := []int{5, 12, 3, 9, 7, 1}
	fmt.Println("ID-uri (in ordinea inelului):", ids)
	fmt.Println("Liderul ales:", RingFlooding(ids))
}
