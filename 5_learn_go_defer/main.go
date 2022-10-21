package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("And Welcome")
	fmt.Println("Hello")
	myDefer()
}
func myDefer() {
	var i int
	for i = 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
