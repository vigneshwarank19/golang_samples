package main

import (
	"fmt"
	"time"
)

type structure struct {
	val int
	str string
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() { // goroutine & asyncrones

	f("direct")

	//	go f("goroutine")

	//	go f("hello peter")

	f("hi vk")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	f("strange")

	go f("iam mad")

	time.Sleep(15 * time.Second)
	fmt.Println("done")

	// how to get struct values

}
