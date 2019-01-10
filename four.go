package main

import "fmt"

func basicFor() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum +=i
	}
	return sum
}

func main() {
	fmt.Println(basicFor())
}