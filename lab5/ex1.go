package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Robot struct{ x, y float64 }

const raza = 20.0
const eps = 1e-6
const maxIteratii = 10000

func distanta(a, b Robot) float64 {
	return math.Hypot(a.x-b.x, a.y-b.y)
}

func citeste(nume string) ([]Robot, error) {
	file, err := os.Open(nume)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return nil, err
	}
	if n < 1 || n > 50 {
		return nil, fmt.Errorf("numarul de roboti trebuie sa fie intre 1 si 50, nu %d", n)
	}

	roboti := make([]Robot, 0, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		x, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		scanner.Scan()
		y, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}
		roboti = append(roboti, Robot{x, y})
	}

	return roboti, scanner.Err()
}

func comprima(roboti []Robot) []Robot {
	activi := []Robot{}
	for _, r := range roboti {
		nou := true
		for _, a := range activi {
			if distanta(r, a) < eps {
				nou = false
				break
			}
		}
		if nou {
			activi = append(activi, r)
		}
	}
	return activi
}

func main() {
	roboti, err := citeste("input.txt")
	if err != nil {
		fmt.Println("Eroare la citire:", err)
		return
	}

	roboti = comprima(roboti)
	fmt.Println("Roboti activi la inceput:", len(roboti))

	iteratii := 0
	for len(roboti) > 1 && iteratii < maxIteratii {
		noi := make([]Robot, len(roboti))
		legat := false
		for i, r := range roboti {
			sumaX, sumaY, nr := 0.0, 0.0, 0
			for j, v := range roboti {
				if i != j && distanta(r, v) < raza {
					sumaX += v.x
					sumaY += v.y
					nr++
				}
			}

			if nr > 0 {
				legat = true
			}
			noi[i] = Robot{(r.x + sumaX) / float64(1+nr), (r.y + sumaY) / float64(1+nr)}
		}

		if !legat {
			break
		}

		iteratii++
		roboti = noi

		inainte := len(roboti)
		roboti = comprima(roboti)
		if len(roboti) != inainte {
			fmt.Printf("Iteratia %d: %d roboti activi\n", iteratii, len(roboti))
		}
	}

	if len(roboti) == 1 {
		fmt.Printf("Rendezvous la (%.4f, %.4f) dupa %d iteratii\n", roboti[0].x, roboti[0].y, iteratii)
	} else {
		fmt.Printf("Nu s-a ajuns la un singur punct (iteratii facute: %d). Pozitii ramase:\n", iteratii)
		for _, r := range roboti {
			fmt.Printf("(%.4f, %.4f)\n", r.x, r.y)
		}
	}
}
