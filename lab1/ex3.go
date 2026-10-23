package main

import "fmt"

func main() {
	sir := [10]int{12, 7, 4, 9, 20, 15, 8, 3, 6, 11}
	pare := []int{}

	for _, x := range sir {
		if x%2 == 0 {
			pare = append(pare, x)
		}
	}

	fmt.Println("Sirul:", sir)
	fmt.Println("Numere pare:", pare)
}
