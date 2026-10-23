package main

import "fmt"

func main() {
	s := []string{"Laboratorul", "de", "sisteme", "distribuite"}

	// schimbam primul cu ultimul, al doilea cu penultimul etc.
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - 1 - i
		temp := s[i]
		s[i] = s[j]
		s[j] = temp
	}

	fmt.Println(s)
}
