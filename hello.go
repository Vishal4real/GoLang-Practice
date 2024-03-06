package main

import (
	"fmt"
)

//this is Single line comment
/* this is multi line comment
hehehehhheeh
*/
func main() {
	//fmt.Println("Hello World")
	var greet string = "hello world"
	var greeet = "And"
	greeting := "bye world"
	fmt.Println(greet, greeet, greeting) // THis will print hello world And bye world

	//Declearing Variables Without inital value

	//var greet --> Error

	var a int
	//assiging value to variable a
	a = 20
	fmt.Println(a)

	// Declearing Multiple
	var x, y, z = 10, "ggg", false
	fmt.Println(x, y, z)
}
