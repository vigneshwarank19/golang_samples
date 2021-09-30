package main

import "fmt"

func main() {
	var a = map[string]int{"peter": 1, "deadman": 2}           // "u"
	b := map[string]string{"peter": "hello", "deadman": "whf"} // "u"
	fmt.Println(a)
	fmt.Println(b)
	var c = make(map[string]int) // "u"
	c["peter"] = 1
	c["wick"] = 2
	fmt.Printf("%v", c)
	d := make(map[string]string) // "u"
	d["sandman"] = " parker"
	d["goblin"] = "peter"
	fmt.Printf("%v \n", d)
	// is some different in empty map
	var e = make(map[string]string)
	var f map[string]string
	fmt.Println(e == nil)
	fmt.Println(f == nil)
	* function
	* map */
	// but values can be any type
	// how to access map element
	fmt.Println(d["goblin"])
	// updating and adding elements in map
	//adding
	d["sockman"] = "spiderman"
	// update
	d["goblin"] = "sandman"
	fmt.Println(d)
	// remove element from map
	delete(d, "goblin")
	fmt.Println(d)
	// checking specific element and value is here or not
	val, ok := d["sandman"]
	fmt.Println(val, ok)
	// if u dont want val omit it use this insted of val - "_"
	_, ok1 := d["goblin"]
	fmt.Println(ok1)

	// Maps are references to hash tables.
	// If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	// map take deep copy
	g := map[string]int{"vk": 1, "kk": 2}
	h := g
	// we change the "h" element
	h["kk"] = 5
	fmt.Println("after changing")
	fmt.Println(g)
	fmt.Println(h)
	// iterating map using "range"
	for k, v := range g {
		fmt.Printf("%v \t %v \n", k, v)
	}
	// Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
	// how to order specifi the map
	i := map[string]int{"vv": 1, "kk": 2, "vk": 3, "kv": 4}
	var j []string
	j = append(j, "kv", "vk", "kk", "vv")
	//Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
	// non-specified order
	for ke, v := range i {
		fmt.Printf("%v : %v \n", ke, v)
	}
	//order specifi
	for _, val := range j {
		fmt.Printf("%v: %v \n", val, i[val])
	}

	/* difference btwn struct and map : struct can't add values if we need but map can. */
}
