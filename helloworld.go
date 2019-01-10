package main

import (
	"fmt"
	"time"
)

func main() {
	go printHelloworld()
	time.Sleep(5*time.Second)
}

func printHelloworld() {
	fmt.Println("hello wordl!")
}