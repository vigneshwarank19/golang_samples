package main

import "fmt"

func main() {
	a := 3
	fmt.Println(recursion(a))
}
func recursion(x int) (y int) {
	if x > 0 {
		y = x * recursion(x-1)

	} else {
		y = 1
	}
	return
}
