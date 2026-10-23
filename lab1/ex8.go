package main

import "fmt"

func main() {
	angajati := map[string]int{
		"Ana Popescu":      1,
		"Mihai Ionescu":    2,
		"Bogdan Radu":      1,
		"Elena Dumitrescu": 3,
		"Cristian Vasile":  2,
		"Ioana Enache":     1,
	}

	departamente := make(map[int][]string)
	for nume, dep := range angajati {
		departamente[dep] = append(departamente[dep], nume)
	}

	for d := 1; d <= len(departamente); d++ {
		lista := ""
		for i, nume := range departamente[d] {
			if i > 0 {
				lista += ", "
			}
			lista += nume
		}
		fmt.Printf("Departamentul %d: [%s]\n", d, lista)
	}
}
