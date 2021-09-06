package main

import "fmt"

func main() {
	var chiken map[string]int
	chiken = map[string]int{}

	chiken["January"] = 50
	chiken["Febuary"] = 40

	fmt.Println("January")
	fmt.Println("Febuary", chiken["January"])

	// var data map[string]int
	// data ["one"] = 1

	// error
	var data map[string]int

	data = map[string]int{}
	data ["one"] = 1

	fmt.Println("data", data["one"])

	var chiken2 = map[string]int{
		"january" : 30,
		"febuary" : 40,
	}
	fmt.Println("Kodok", chiken2["january"])


	var ary = map[string]string{
		"Belut laut" : "Kodok",
		"Ikan Laut"	: "penyu",
	}
	ary["Hiu laut"] = "Buaya"

	fmt.Println("KODOK", ary["Belut laut"])
	fmt.Println("Lumba lumba", ary["Hiu Laut"])
	fmt.Println(ary)

	delete(ary, "Ikan Laut")

	fmt.Println(ary)
}