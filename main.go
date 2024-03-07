package main

import "fmt"

func main() {
	for i := 0; i <= 10; i = i + 1 {

		switch i {
		case 2:
			i = i * 3
			fmt.Println(i)
		case 10:
			fmt.Println(("終了"))
		default:
			fmt.Println(i)
		}

	}
}
