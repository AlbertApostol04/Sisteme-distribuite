package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	cuvinte := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linie := scanner.Text()

		cuvant := ""
		for _, c := range linie {
			if c == ' ' {
				if cuvant != "" {
					cuvinte = append(cuvinte, cuvant)
					cuvant = ""
				}
			} else {
				cuvant += string(c)
			}
		}
		if cuvant != "" {
			cuvinte = append(cuvinte, cuvant)
		}
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Eroare la citire: ", err)
		return
	}

	aparitii := make(map[string]int)
	for _, c := range cuvinte {
		aparitii[c]++
	}

	for cuvant, nr := range aparitii {
		fmt.Printf("%s: %d\n", cuvant, nr)
	}
}
