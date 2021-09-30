/* the golong have only loop is for */
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("        ")
	fmt.Println("        ")
	// nested loop
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	for j := 0; j < len(a); j++ {
		for k := 0; k < len(b); k++ {
			fmt.Println(j, "and", k)
		}

	}
	// range
	e := []string{"hi", "hello", "hmm"} // - "u"
	for c, d := range e {
		fmt.Printf("%v \t %v \n", c, d)
	}

	/* if we don't want index or value just put that place at "_" it will omit that  */
	f := []string{"rr", "kk", "vv"} // - "u"
	for _, h := range f {
		fmt.Printf("%v\n", h)
	}
	//  we can also use "for loop" this type (we don't need to put semicolons at infront of g and backward of g like other language)
	g := 10
	for g > 0 {
		fmt.Println("hunt")
		g--
	}

}
