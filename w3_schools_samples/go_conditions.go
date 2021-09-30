package main

import "fmt"

func main() {
	a := 5
	if a > 10 {
		fmt.Println("it is large")
		if a < 30 {
			fmt.Println(" it is a small num ")
		} else {
			fmt.Println(" it is not a number ")
		}
	} else if a > 20 { // "u" , note every after if last "}" strickly for else that's only run otherwise throw error, if we separate "}" and else it will give throws
		fmt.Println("it is medium")
	} else {
		fmt.Println("it is small")
	}
	// difference bwn other lang vs golang (golang condition and switch don't have paramthies"()",switch type only same type in golang expression and value type and also we don't put a expression golang can run but other lang we must give expression or value or variable and even we can put a expression in cases it will run in golang )
}
