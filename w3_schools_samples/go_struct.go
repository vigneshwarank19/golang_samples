package main

import "fmt"

// struct is like a class (golang don't have class)
type supers struct { // "u" , "sy",struct can store multi-different data-types (struct like a blueprint)
	spider string
	iron   string
	dhoni  int
}
type legends struct {
	dhoni string
	virat string
	rohit string
	no1   int
	no2   int
	no3   int
}
type eng struct {
	what string
	who  string
	when string
}

// sruct is user-define-data-type.
// struct is like a array store a multi-value in single value but array can't accept different data-types
// we have two ways to assaign a value into struct
func main() {

	var heros supers // struct use like a data-type eg:int,string , inside of the structure ever one is call as "member"
	heros.dhoni = 07 // this is one way to "assaign a values into struct"
	heros.spider = "spider_man"
	heros.iron = "iron_man"
	str(heros)
	var ptr *supers  // we can do struct as a pointer
	fmt.Println(ptr) // pointer default value is nil
	var players legends
	players.dhoni = "finisher"
	players.virat = "classic"
	players.rohit = "hitter"
	players.no1 = 07
	players.no2 = 18
	players.no3 = 13
	pyrs(&players)                          // we sent a address to be a function
	var know eng = eng{"iam", "iam", "now"} // this is another way to assaign a value into the struct
	fmt.Println(know)
}

// "u" , we can pass a struct as a function argument in normal way.
func str(sups supers) {
	fmt.Println(sups.spider) // if we want to access a value of struct use "."
	fmt.Println(sups.iron)
	fmt.Println(sups.dhoni)
}
/* we can do pointer to struct*/
func pyrs(play *legends) { //"u" , pointer getinto that address and get the value of the address
	fmt.Println(play.dhoni)
	fmt.Println(play.virat)
	fmt.Println(play.rohit)
	fmt.Println(play.no1)
	fmt.Println(play.no2)
	fmt.Println(play.no3)
}
