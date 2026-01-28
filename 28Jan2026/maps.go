package main

import "fmt"

func changeC(m map[string]int) {
	// m["c"] = 300

	// reassigning map to a new map
	m = make(map[string]int)
	m["c"] = 300
	fmt.Println("Map m inside changeC function:", m)
}

func main() {
	// var m map[string]int // nil map

	// if m == nil {
	// 	fmt.Println("Map is nil!", m)
	// 	fmt.Println("Value for \"x\":", m["x"]) // zero value for int type - non-existing key
	// 	m["x"] = 10                             // panic: assignment to entry in nil map
	// }

	mp := map[string]int{} // empty map
	fmt.Println("Empty Map:", mp)
	fmt.Println("Value for \"x\":", mp["x"]) // zero value for int type - non-existing key
	mp["x"] = 10
	fmt.Println("Map after adding key-value pair:", mp)

	m := make(map[string]int) // empty map using make - we can also specify capacity using make
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
	fmt.Println("Map m:", m)

	// loop through map
	for key, value := range m {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// delete key "b"
	delete(m, "b")
	fmt.Println("Map m after deleting key \"b\":", m)

	// check if key "c" exists
	value, ok := m["c"]
	if ok {
		fmt.Println("Key \"c\" exists with value:", value)
	} else {
		fmt.Println("Key \"c\" does not exist")
	}

	// check if key "b" exists
	value, ok = m["b"]
	if ok {
		fmt.Println("Key \"b\" exists with value:", value)
	} else {
		fmt.Println("Key \"b\" does not exist")
	}

	changeC(m)
	fmt.Println("Map m after changeC function call:", m)
}
