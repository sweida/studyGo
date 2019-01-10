package main

import "fmt"

func main() {
	var maxNum int = 10

	for i := 1; i < maxNum; i++ {
		for j :=1; j <= i; j++ {
			fmt.Printf("%v * %v = %v\t", i , j, i*j)
		}
		// 换行
		println()
	}
}