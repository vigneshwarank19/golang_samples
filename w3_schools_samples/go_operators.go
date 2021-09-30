package main

import "fmt"

func main() {
	var a int = 5
	b := a + 5
	fmt.Println(b) // add " + "

	c := b - 1
	fmt.Println(c) //sub

	d := c * 2
	fmt.Println(d) //mul

	e := d / 2
	fmt.Println(e) //div

	f := d % 2
	fmt.Println(f) //mod

	g := a
	a--            // "u" golong doesn't wrk in pre-incre and pre- dec it will give error it's a bug or something
	fmt.Println(g) //dec

	h := a
	a++
	fmt.Println(h) //inc

	i := 3
	i -= 1
	fmt.Println(i) // assignment opr

	add := 1 // comparision opr
	sub := 2
	mul := 3
	div := 1
	mod := 5
	fmt.Println(add == div)
	fmt.Println(sub > mul)
	fmt.Println(sub < add)
	fmt.Println(mod >= add)
	fmt.Println(mod <= mul)
	fmt.Println(add != sub)

	// logic operator
	fmt.Println("the add opr", (add == div && mod > div))
	fmt.Println("the or opr", (add == div || mod < div))
	fmt.Println("the not opr", !(add == div && mod > div))

	// bitwise operator
	fmt.Println(add & sub)
	fmt.Printf("%b", add&sub)
	fmt.Printf("%b", add|sub)
	fmt.Printf("%b", add^sub)
	//	fmt.Println("%b", add<<sub)
	/* "u" - it's give error becz only unsigned value be assign .
	it's clear in 1.13 version in golong check sTACK OVER FLOW */
	//	fmt.Printf("%b", add>>sub) - "u"
	const (
		cn = 3
	)
	cn1 := 1 << cn
	fmt.Printf("%b \n", cn1)
	/*relation opertor and logic operator give boolen value only*/

}
