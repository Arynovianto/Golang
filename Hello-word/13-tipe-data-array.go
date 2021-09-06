package main

import "fmt"

func main() {
	var name = [3]string{
		"Muhamad",
		"Ary",
		"Novianto",
	}

	fmt.Println(name)
	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var value = [4] int {
		10,
		20,
		30,
	}
	fmt.Println(value)
	fmt.Println(value[0])
	fmt.Println(value[1])
	fmt.Println(value[2])

	fmt.Println(len(value))
	fmt.Println(len(name))
}