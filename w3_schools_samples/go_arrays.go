package main

import "fmt"

var a int = 5
var arr1 = [3]string{"rko", "cena", "bbm"} // len defined ("syntax-sy") , (element - ele)
var arr2 = [...]int{1, 2, 3, 4, 5, 6, 7}   // inferred
func main() {
	arr3 := [2]string{"ff1", "ff2"} // "u" , "sy"
	arr4 := [...]int{1, 2, 3}       // "u" , "sy"
	fmt.Println(a)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr1)
	arr1[2] = "jeff" // "u" , var will not hoisting , changing ele "sy" into a array
	arr3[1] = "ff3"  // "u"  , inferred also not hoisting the value
	fmt.Println(arr1)
	fmt.Println(arr3)
	fmt.Println(arr2)
	fmt.Println(arr1[1]) // "u" , accesing a ele into a array
	a = 6                // "u"
	fmt.Println(a)
	arr5 := [5]int{1, 2}     // "u" , half intial
	arr6 := [5]int{}         // "u" , empty array
	arr7 := [5]string{"hii"} // "u" , "string" - default value is empty - ("")
	arr8 := [5]string{}      // "u"
	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)
	arr9 := [5]int{2: 20, 4: 40} // "u" , we intialize only specific  element index (ind - 2 - ele - 20)
	fmt.Println(arr9)
	/* if we want to find the length we use below this (len())*/
	arr10 := [...]int{1, 2, 3}
	var arr11 = [3]string{"hii", "hiii", "hiii"}
	fmt.Println(len(arr10))
	fmt.Println(len(arr11))
	var arr15 = []int{1, 2, 3} // "u" another type
	fmt.Println(arr15)

	var mu = [2][3]int{{1, 2, 3}, {0, 1, 3}} // "u" - multi-dimension array
	fmt.Println(mu)

}
