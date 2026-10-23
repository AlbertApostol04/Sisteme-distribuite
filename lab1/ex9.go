package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	arr := [6]int{2, 7, 11, 8, 4, 9}
	fmt.Println("Array:", arr)

	fmt.Print("s = ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	s := 0
	for _, c := range text {
		s = s*10 + int(c-'0')
	}

	vazute := make(map[int][]int)

	for j, val := range arr {

		for _, i := range vazute[s-val] {
			fmt.Println(i, j)
		}
		vazute[val] = append(vazute[val], j)
	}
}
