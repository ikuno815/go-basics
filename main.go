package main

import "fmt"

func main() {
	fmt.Println("hola el mundo")
	sl := []int{1, 2, 3}
	if len(sl) > 2 {
		fmt.Println("unreachable code")
	}
}
