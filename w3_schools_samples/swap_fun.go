package main

import "fmt"

func swap(x, y int) (a, b int) {
	b = x
	a = y
	return
}

func main() {
	fmt.Println(swap(5, 6))
}
