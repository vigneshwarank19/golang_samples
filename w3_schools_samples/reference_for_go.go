package main // every program must need (package main)
/* we dont use ";" because "golong" take line break ("enter click") as ";" , if we pur two var or const or something else
in same line it would give erro if we over come that we will use ";" between of them it will execute'                                                                           */
import "fmt"           // libray (fmt)
const thr bool = false //"u" , const like a var we use inside and outside blocks.
const frt = 11.2       // inferred
func main() {
	// if u got problem first check whole syntax , it would be the issue
	var f int = 5
	var s = 6 //  type is inferred (compiler will decided value basis)
	t := 7    // ":=" it's like a variable but compiler will decide what type as value basis
	/* var can inside and outside the fun (main), declaration and assign can separate
	   ":" can only inside ,can't separate val and var
	*/
	var fo string // unique - "u"
	fo = "work it"
	var fi, si, se int // "u"
	fi, si, se = 1, 2, 3
	var ei, ni = 5, "yummy" // "u" , if we haven't mention the type we woundn't separate var and val it will give error
	var (
		te int    = 3
		el string = "horrible"
		tw string = "thoughts"
	) // "u" , block type
	const (
		fif int = 1
		six     = "hunt"
	) // "u"
	var sevn int = 500
	var eig uint = 4500 /* int have two typs signed and unsigned unsigned only store positive val but signed store both
	"+ and - val" if u dont mention signed or unsigned it will take signed and also it take "bit "
	depends on platform(system)and also we mention what "bit" : (eg :int64 = 1373892 ) and
	 we put wrong bit it will overflow and also we dont mention float "bit"  it will take ""float64*
	 it is not like "int" not depend on platform */
	result := f + s + t
	fmt.Println("the value of result is : ", result, " ", fo)
	fmt.Print(thr, "\n", frt) // "u",print same line,"output types" - "print fun" . we also use " escape sequence " - ""\n""
	fmt.Print("variant")
	fmt.Println("loki")
	fmt.Printf("value is : %v \n type is : %v \n", f, f) // "u" , format specifier or formatting verbs
	fmt.Printf("value is : %v \n type is : %v", s, s)    /* "u", formatting verbs have lot of types so if u stuck something like
	   this ("%" if it anything "%" between the double quotes is a formatting verb ) */
	fmt.Println(fi, si, se)
	fmt.Println(te, el, tw)
	fmt.Println(sevn)
	fmt.Println(eig)
	fmt.Println(ei, ni)
	va, vaa, vaaa, vaaaa := "hi", 2, "hell", 3.45
	fmt.Println(va, vaa, vaaa, vaaaa) // we can declare same line in multi values and multi variables and those can different types
	const v int = 10                  // we can't declare separate declaration and value
	fmt.Println(v)
}
