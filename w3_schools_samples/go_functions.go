package main

import "fmt"

func vk(x int, y int) int { // - "u"
	a := x + y
	return a
}

//named return value
/*we can store a retrun value in  a variable if we did this we need to just give return
it will give the return var value */
/* func kk(x int, y string) (c int, d string) { // - "u"
	c = x
	d = y
	return
}
*/
func vic(x int, y string) (result int, txt1 string) { // we can get a multiple return types in golang
	result = x
	txt1 = y
	return
}

// "u" this is also different type , we mention only return types is "int" but we didn't mention which one would come (x or y)
func kv(x, y int) int {
	if x < 5 {
		return x
	} else {
		return y
	}
}

// "u" this another methode "variable use like a function (anonimous function)"
var ty = func(a, b int) (c, d int) {
	c = a
	d = b
	return
}

func main() {
	fmt.Println(vk(2, 3),
		vk(5, 6),
		vk(7, 8))
	a, b := vic(5, "Hello") // fisrt you need to give function call and print will give separate if doesn't do it it give error
	fmt.Println(a, b)
	fmt.Println(kv(10, 40))
	fmt.Println(ty(5, 6))

}
