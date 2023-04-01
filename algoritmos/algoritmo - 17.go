package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n [10]int
	var soma int

	// randomizer
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(100)
	}

	// calcula a soma dos elementos pares
	for i, num := range n {
		if i%2 == 0 {
			soma += num
		}
	}

	fmt.Println(soma)
}
