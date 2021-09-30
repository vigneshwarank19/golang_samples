package main

import "fmt"

func main() {
	/*slice - slice is like a array but arrays are fixed length we can't add lot but slice is dynamic len
	we add what length we want  */
	/*difference btwn array and slice : array size are fixed we can't grow or shrink but slice can grow and shrink*/
	/*if we slice a array the slice isn't hold a value or copy the value .
	it will point the original array(deep copy) if we change anything in slice it will change also original array*/
	// 3 way to create a slice
	//way 1 - "u"
	sil1 := []int{1, 2} // normal array method
	fmt.Println(sil1)
	fmt.Println("the length of array : ", len(sil1))   // number of element
	fmt.Println("the capacity of array : ", cap(sil1)) // capacity of element
	// way 2 - "u"
	arr12 := [5]int{1, 2, 3, 4, 5}
	sil2 := arr12[1:4]   // "u" , it will take reference or shallow copy this
	fmt.Println(sil2[0]) // "u" , slice access ele
	sil2[0] = 10
	fmt.Println(sil2)
	fmt.Println(arr12)
	// way 3 - "u"
	// we also create a slice with " make() "
	sil3 := make([]int, 5, 10) // "u" , 2 ele is a len and 3 ele is a capacity and cap is optional
	fmt.Println(sil3)
	fmt.Println(len(sil3))
	fmt.Println(cap(sil3))
	sil4 := make([]string, 2) // "u"
	fmt.Println(sil4)
	fmt.Println(len(sil4))
	fmt.Println(cap(sil4))
	// append (adding a ele into the slice) , it will add last
	sil5 := []int{1, 2, 3}
	sil5 = append(sil5, 10, 20, 30) // "u"
	fmt.Println(sil5)
	// compine 2 slice to make one slice using append
	sil6 := []int{1, 2, 3}
	sil7 := []int{4, 5, 6}
	sil8 := append(sil6, sil7...) // "u" , we must use the "..." after the last slice
	fmt.Println(sil8)
	// change the len of slice
	arr13 := [5]int{1, 2, 3, 4, 5}
	sil9 := arr13[1:4] // slice
	fmt.Println(sil9)
	sil9 = arr13[1:3] // re-size the slice , don't use colon here , if we change already exit slice
	fmt.Println(sil9)
	sil9 = append(sil9, 20, 30) // change the len use append
	fmt.Println(sil9)
	// memory efficiency , "u" , " copy () "
	arr14 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} // original array
	fmt.Println(arr14)
	fmt.Println(len(arr14))
	fmt.Println(cap(arr14))
	//create copy with only needed array
	sil10 := arr14[:len(arr14)-10]
	fmt.Println(sil10)
	cp1 := make([]int, len(sil10))
	fmt.Println(cp1)
	copy(cp1, sil10) // destination and source
	fmt.Println(cp1)
	cp2 := cp1
	fmt.Printf("%v", cp2)

}
