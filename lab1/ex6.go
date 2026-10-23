package main

import "fmt"

func main() {
	inaltimi := map[string]int{
		"Andrei": 190,
		"Maria":  165,
		"Ion":    182,
		"Elena":  172,
		"George": 178,
		"Ana":    165,
	}

	nume := []string{}
	for n := range inaltimi {
		nume = append(nume, n)
	}

	for i := 0; i < len(nume)-1; i++ {
		for j := 0; j < len(nume)-1-i; j++ {
			if inaltimi[nume[j]] < inaltimi[nume[j+1]] {
				temp := nume[j]
				nume[j] = nume[j+1]
				nume[j+1] = temp
			}
		}
	}

	for _, n := range nume {
		fmt.Println(n)
	}
}
