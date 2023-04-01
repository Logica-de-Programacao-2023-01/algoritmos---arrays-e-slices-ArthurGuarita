package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n [10]int
	var s []int

	for i := 0; i < len(n); i++ { // randomizar elementos do array
		n[i] = rand.Intn(100)
	}

	for _, num := range n {
		if num%2 == 0 {
			s = append(s, num)
		}
	}
	fmt.Println(s)
}
