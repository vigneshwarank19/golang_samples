package main

import "fmt"

/*	Concept & Description
1.Go - Array of pointers
You can define arrays to hold a number of pointers.

2.Go - Pointer to pointer
Go allows you to have pointer on a pointer and so on.

3.Passing pointers to functions in Go
Passing an argument by reference or by address both enable the passed argument to be changed in the calling function by the called function.*/

func main() {
	var a int = 1   // when we create a variable that time system give a memory address to that variable , address like a hexa-decimal value
	fmt.Println(&a) // if we want a address of the variable we type "&variable_name"
	// pointer can store a address of the variable
	var ptr *int // "u" ,"syn" , declaration
	ptr = &a     // assign a address of the variable or another name is "reffering"
	fmt.Println(ptr)
	fmt.Println(*ptr) /*if we want a value from the address ,"syn"-"*pointer_name" or another name is "dereffering"*/
	var pptr **int    // if u need to refer a pointer to pointer just include a how many times u point that variable
	pptr = &ptr       // put a "&" sysmbol must
	fmt.Println(pptr)
	fmt.Println(**pptr)
	// pointer give a nil value it haven't know the address or haven't know the proper address
	var ppptr *int
	fmt.Println(ppptr) // nil
	// we can hold a pointers in a array
	// we can passing a arugument as a pointer ("passing pointer to a function in a golang")
	// we can hold a pointer to pointer
}
