package main

import "fmt"

func votMajoritar(voturi []float64, adj [][]int) (consens float64, iteratii int) {
	n := len(voturi)
	x := make([]float64, n)
	copy(x, voturi)

	const eps = 1e-9

	for iteratii = 0; iteratii < 100000; iteratii++ {
		lo, hi := x[0], x[0]
		for _, v := range x {
			if v < lo {
				lo = v
			}
			if v > hi {
				hi = v
			}
		}
		if hi-lo < eps {
			break
		}

		inbox := make([][]float64, n)
		for i := 0; i < n; i++ {
			for _, j := range adj[i] {
				inbox[j] = append(inbox[j], x[i])
			}
		}

		nou := make([]float64, n)
		for i := 0; i < n; i++ {
			suma := x[i]
			for _, v := range inbox[i] {
				suma += v
			}
			nou[i] = suma / float64(len(inbox[i])+1)
		}
		x = nou
	}

	return x[0], iteratii
}

func ruleaza(voturi []float64, adj [][]int) {
	suma := 0.0
	for _, v := range voturi {
		suma += v
	}
	media := suma / float64(len(voturi))

	consens, it := votMajoritar(voturi, adj)

	fmt.Println("Voturi:", voturi)
	fmt.Printf("Media reala a voturilor: %f\n", media)
	fmt.Printf("Valoare de consens:      %f (dupa %d iteratii)\n", consens, it)

	if consens > 0.5+1e-6 {
		fmt.Println("Rezultat: majoritatea a votat 1")
	} else if consens < 0.5-1e-6 {
		fmt.Println("Rezultat: majoritatea a votat 0")
	} else {
		fmt.Println("Rezultat: egalitate")
	}
	fmt.Println()
}

func main() {
	n := 9
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = []int{(i + 1) % n, (i + n - 1) % n, (i + 2) % n, (i + n - 2) % n}
	}

	ruleaza([]float64{1, 0, 1, 1, 0, 1, 0, 1, 1}, adj)
	ruleaza([]float64{0, 0, 1, 0, 1, 0, 0, 1, 0}, adj)
}
