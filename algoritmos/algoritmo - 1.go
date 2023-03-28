package main

import "fmt"

func main() {
	var numeros = [3]int{1, 2, 3}
	var result int
	result = 0
	for _, num := range numeros {
		result += num
	}
	fmt.Println(result)
}
