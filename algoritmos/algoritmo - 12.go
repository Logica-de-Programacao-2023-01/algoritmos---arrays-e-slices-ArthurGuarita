package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var array [5]int
	var multiplo int
	for i := 0; i < len(array); i++ {
		array[i] = rand.Intn(100)
	}
	var slime []int
	for _, num := range array {
		multiplo = num % 3
		if multiplo == 0 {
			slime = append(slime, num)
		} else {
			continue
		}
	}
	fmt.Print(slime)
}
