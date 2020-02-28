package main

import "fmt"

func main() {
	m := map[string]int{}
	m["a"] = 1
	m["b"] = 2
	m1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(m, m1)

	// Iterating over a map.
	for k, v := range m {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	// Accessing elements in a map.
	ma := m["a"]
	fmt.Printf("ma=%v\n", ma)
	mb, ok := m["b"]
	fmt.Printf("mb=%v ok=%v\n", mb, ok)
	// Accesing elements that doesn't exist will give you the default value.
	mc := m["c"]
	fmt.Printf("mc=%v\n", mc)
	md, ok := m["d"]
	fmt.Printf("md=%v ok=%v\n", md, ok)

}
