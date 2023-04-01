package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite o valor de n: ")
	fmt.Scan(&n)

	var primo = make([]int, 0)
	for i := 2; n > len(primo); i++ {
		ePrimo := true
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				ePrimo = false
				break
			}
		}
		if ePrimo {
			primo = append(primo, i)

		}
	}
	fmt.Println(primo)
}
