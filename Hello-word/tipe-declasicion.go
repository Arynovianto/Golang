package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var NoKtpAry NoKtp = "20190801334"
	var MarriedStatus Married = true
	fmt.Println(NoKtpAry)
	fmt.Println(MarriedStatus)
}