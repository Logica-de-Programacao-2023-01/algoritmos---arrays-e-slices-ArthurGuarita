package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n = make([]int, 10)
	for i := 0; i < len(n); i++ {
		n[i] = rand.Intn(100)
	}
	var min, max int = n[0], n[0]
	for _, num := range n {
		if num < min {
			min = num
		} else if num > max {
			max = num
		}
	}
	fmt.Println("Valor máximo é", max, "e o valor mínimo é", min)
}
