package main

import "fmt"

func main() {
	var r int
	fmt.Printf("enter the limit of pyramid")
	fmt.Scanln(&r)
	for b := 1; b <= r; b++ {


	for a := 1; a <=b; a++ {
		fmt.Print("* ")

	}
		fmt.Printf(" \n")
}
}