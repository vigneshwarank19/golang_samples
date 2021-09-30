package main

import "fmt"

/* interface is a another data-type in golang , it will give unimplement methodes (methode signature).
unimplement methode like (methode name and arguments and that's retyrn type only it will give).
The struct data type implements these interfaces to have method definitions
for the method signature of the interfaces.*/
/*other oops programming have "class". but golang don't have class */

type inter interface {
	meth1() int
	// meth2 () string
}
type struc struct {
	a int
	b int
}

/* type struc2 struct{
	c string
	d string
 } */

func (s struc) meth1() int {
	return s.a + s.b
}

/* func (s struc)meth2()string{
	  return struc2.c + struc2.d
  } */
func funi(i inter) int {
	return i.meth1()
}
func main() {
	val := struc{a: 4, b: 5}
	// fmt.Println(val)
	fmt.Println(funi(val))

}
