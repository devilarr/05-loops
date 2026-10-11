package main

import "fmt"
fork test
func main(){
	age:= 20

	if age >= 18{
		fmt.Println("You are an adult")
	}	 else if age >= 12{	
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("You are a minor")
	}
}
