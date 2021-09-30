package main

import "fmt"

// Type conversion is a way to convert a variable from one data type to another data type.
// "sy" - type_name(expression)
/* Type conversion - some times we can't convert it give "invalid conver"
,"may be types have also "precedance" that's why we face some problem" */
func main() {
	var a int = 5
	var b int = 6
	var c float64
	// c = a + b - it will throw error
	c = float64(a) + float64(b)
	fmt.Println(c)
	fmt.Printf("%T %T %T \n", a, b, c)
	var d int = 5
	var e int = 6
	var f float32
	f = float32(d) + float32(e)
	fmt.Println(f)

}
