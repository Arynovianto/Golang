package main

import (
	"fmt"
)

func main() {
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	// int8 cangkupan datanya nya -128 sampai 127
	// int16 cangkupan datanya -32768 sampai 32767
	// int32 cangkupan datanya -2147483648 sampai 2147483647
	// int64 cangkupan datanya -9223372036854775808 sampai 9223372036854775807

	// unt8 cangkupan datanya 0 sampai 255
	// unt16 cangkupan datanya 0 sampai 65535
	// unt32 cangkupan datanya 0 sampai 4294967295
	// unt64 cangkupan datanya 0 sampai 18446744073709551615


	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var nama string = "Muhamad Ary Novianto"
	var e byte = nama[0]
	var eString = string(e)

	fmt.Println(nama)
	fmt.Println(eString)
}