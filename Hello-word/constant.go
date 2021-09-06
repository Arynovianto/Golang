package main

import "fmt"

func main() {
	const (

		// gAK PAKE TULISAN string juga bisa
		fistname string= "Muhamad "
		lastname = "Ary "
		endname  = "Novianto"
		number   = 1000
	)
	fmt.Println(fistname+lastname+endname)
	fmt.Println(lastname+endname)
	fmt.Println(lastname)
	fmt.Println(endname)
	fmt.Println(fistname)

	fmt.Println(number)

	// // Error
	// fmt.Println(fistname+number)
}