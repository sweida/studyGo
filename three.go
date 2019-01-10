package main

import (
	"fmt"
)

func main() {

	// Find the quotient, reminder and module of x / y.
	div := func(x, y int) (quotient float64, reminder, modulo int) {
		quotient = float64(x) / float64(y)
		reminder = x % y
		modulo = 3
		return
	}

	q, r, m := div(-21, 4)
	fmt.Printf("三个值为%v %v %v", q, r, m)
	if q != -5.25 {
		fmt.Println("求商错误")
	}
	if r != -1 {
		fmt.Println("求余数错误")
	}
	if m != 3 {
		fmt.Println("求什么错误")
	}

}