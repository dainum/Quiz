package main

import "fmt"

func main() {
	m := map[string]int{"x": 10, "y": 20}
	//m2 := m
	m2 := make(map[string]int)
	for k, v := range m {
		m2[k] = v
	}

	m["x"] = 20
	m2["y"] = 0
	fmt.Println(m["x"], m["y"])
	fmt.Println(m2["x"], m2["y"])
}
