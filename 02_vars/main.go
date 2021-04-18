package main

import "fmt"
var name = "Ryan"

func main(){
	var age = 32
	var isAwake = true
	const isImutable = true
	shorthand := "uses := to declare. no const or var. Must be declared inside a function"
	shoeSize := 12.5
	double, declare := "one line", "two variables"
	fmt.Println(name, age)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", shoeSize)
	fmt.Println(isAwake)
	isAwake = false 
	fmt.Println(isAwake)
	fmt.Println(shorthand)
	fmt.Println(double, declare)
}