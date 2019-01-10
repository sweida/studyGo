package main

import (
	"fmt"
)

func main() {
	add := func(x, y int) int {
		return x + y
	}

	sub := func(x, y int) int {
		return x - y
	}

	result := sub(add(1, 2), 3) 
	fmt.Println(result)

	if result != 0 {
		fmt.Println("代码错误")
	} else {
		fmt.Println("成功！")
	}
}
