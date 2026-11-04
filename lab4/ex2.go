package main

import (
	"fmt"
	"sync"
)

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0

	for start := 0; start < n; start++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[start] = 0
		coada := []int{start}

		for len(coada) > 0 {
			u := coada[0]
			coada = coada[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					coada = append(coada, v)
				}
			}
		}

		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

func FloodingGraf(ids []int, adj [][]int) int {
	n := len(ids)
	runde := diametru(adj)
	if n == 0 || runde < 0 {
		return -1
	}

	outCh := make([][]chan int, n)
	inCh := make([][]chan int, n)
	for i := 0; i < n; i++ {
		for _, j := range adj[i] {
			c := make(chan int, runde+1)
			outCh[i] = append(outCh[i], c)
			inCh[j] = append(inCh[j], c)
		}
	}

	leaders := make([]bool, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			maxID := ids[i]

			for r := 0; r < runde; r++ {
				for _, c := range outCh[i] {
					c <- maxID
				}

				for _, c := range inCh[i] {
					v := <-c
					if v > maxID {
						maxID = v
					}
				}
			}

			leaders[i] = maxID == ids[i]
		}(i)
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
	ids := []int{4, 9, 2, 7, 5, 1}

	adj := [][]int{
		{1},
		{2},
		{3},
		{4, 1},
		{5},
		{0},
	}

	fmt.Println("ID-uri:", ids)
	fmt.Println("Diametrul grafului:", diametru(adj))
	fmt.Println("Liderul ales:", FloodingGraf(ids, adj))
}
