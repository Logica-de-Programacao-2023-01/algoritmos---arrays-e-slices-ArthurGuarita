package main

import "fmt"

func main() {
	var nums = [5]int{1, 2, 3, 4, 5}
	for _, num := range nums {
		fmt.Println(num)
	}
}
