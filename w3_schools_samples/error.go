package main

import (
	"errors"
	"fmt"
)

func d(v int) (int, int, error) {
	return v + 1, 1, errors.New("hello peter")
}
func main() {
	a, b, c := d(1)
	fmt.Println(a, b, c)
}
