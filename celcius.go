package main

import "fmt"

func main(){
	fmt.Printf("enter the temperature for conversion")
	var f,c float32
	fmt.Scanln(&f)
	c=(f-32)*5/9
	fmt.Printf("coversiuon of %f f is %f c",f,c )

}
