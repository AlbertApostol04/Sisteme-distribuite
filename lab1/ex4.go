package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := []float32{1.5, 2.5, 3.5, 4.5, 5.5}
	fmt.Println("Slice:", s)

	fmt.Print("Index de eliminat: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	idx := 0
	for _, c := range text {
		idx = idx*10 + int(c-'0')
	}

	if idx < 0 || idx >= len(s) {
		fmt.Println("Index invalid")
		return
	}

	s = append(s[:idx], s[idx+1:]...)
	fmt.Println("Rezultat:", s)
}
