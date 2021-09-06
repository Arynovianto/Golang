package main

import "fmt"

func main() {
	var bulan = [...]string {
		"january",
		"febuary",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	fmt.Println(bulan)

	var slice1 = bulan [4:7]

	fmt.Println(slice1)

	fmt.Println(len(slice1))

	// bulan[5] = "diubah"
	// fmt.Println(slice1)
  // ----------------

	// slice1[0] = "mei lagi"
	// fmt.Println(slice1)

	var slice2 = bulan[2:4]
  var slice2 = bulan[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "ary")
	fmt.Println(slice3)
	
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(bulan)
	
	var newslice = make([]string, 2, 5)
	newslice[0] = "Ary"
	newslice[1] = "Novianto"

	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))
}