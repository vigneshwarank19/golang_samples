package main

import "fmt"

func main() {
	a := "hello"
	switch a {
	case "hi":
		fmt.Println("hi")
	case "hmm":
		fmt.Println("hmm")
	case "hello":
		fmt.Println("hello")
	default:
		fmt.Println("none")

	}
	/* we don't use break statement inbetween the case ,
	All the case values should have the same type as the switch expression.
	 Otherwise, the compiler will raise an error: ,the case and switch expression should be same */

	// multi case
	b := 5
	switch b {
	case 1, 2:
		fmt.Println("value is 1 and 2")
	case 3, 4:
		fmt.Println("value is 3 and 4")
	case 5, 6:
		fmt.Println("value is 6 and 5")
	default:
		fmt.Println("none")
	}
	/* multi case - if we want we put two or more value can put it */

	c := 5
	switch c {
	case c > 2:
		fmt.Println("hell")
	case c < 1:
		fmt.Println("what")
	}

}
