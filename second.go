package main

import (
	"fmt"
)

func SumAdd(x, y int) int {
	return x + y
}

func main() {
	for i := 0; i < 10; i++ {
		result := SumAdd(i, i)
		fmt.Println(result)
	}
}

