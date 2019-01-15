package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("time break 2 second")
	time.Sleep(2 * time.Second)
	fmt.Printf("....")
	time.Sleep(2 * time.Second)
	fmt.Printf(".....")
	time.Sleep(2 * time.Second)
	fmt.Printf("@-@ \n")
	time.Sleep(2 * time.Second)
	fmt.Printf("sucessfully compiled ")

}
