package main

import "fmt"

func main(){
	fmt.Print("enter the two number for addition")
	var a,b int
	var sum int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	sum=a+b
	fmt.Print(sum)
}
