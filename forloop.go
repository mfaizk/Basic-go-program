package main

import "fmt"

func main(){
	var a,b int
	fmt.Print("enter the number for table")
	fmt.Scanln(&a)
	for i:=0; i<=10; i++ {
	b=a*i

	println(b)
}}