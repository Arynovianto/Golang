package main

import "fmt"

func main() {
	var NilaiUjian 		= 75
	var Absensi 			= 75
	
	
	var LulusUjian		= NilaiUjian >= 75
	var LulusAbsensi	= Absensi >= 75 

	fmt.Println(LulusUjian)
	fmt.Println(LulusAbsensi)

	var Lulus 				= LulusUjian && LulusAbsensi

	fmt.Println(Lulus)
}