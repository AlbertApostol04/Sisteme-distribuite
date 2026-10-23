package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	n := 0
	for _, c := range text {
		n = n*10 + int(c-'0')
	}

	prim := true
	if n < 2 {
		prim = false
	}

	for i := n / 2; i > 1; i-- {
		if n%i == 0 {
			prim = false
		}
	}

	if prim {
		fmt.Println("e prim")
	} else {
		fmt.Println("nu e prim")
	}
}
