package main

import (
	"fmt"
)

// this is Single line comment
/* this is multi line comment
hehehehhheeh
*/
func main() {
	//fmt.Println("Hello World")
	/*var greet string = "hello world"
	var greeet = "And"
	greeting := "bye world"
	fmt.Println(greet, greeet, greeting) // This will print hello world And bye world

	//Declearing Variables Without inital value

	//var greet --> Error

	var a int
	//assiging value to variable a
	a = 20
	fmt.Println(a)

	// Declearing Multiple
	var x, y, z = 10, "ggg", false
	fmt.Println(x, y, z)

	// Constants in Go
	// two types of constant Typed and Untyped Constant
	const HEllO = 1234         // untyped Constant
	const HEHE string = "4567" // Typed Constant

	// Declearing Multiple Constant Together \

	const (
		hi        = 12
		ab string = "av"
		tf        = true
	)
	fmt.Println(hi, tf)*/
	//foutput()
	verb()
}

func foutput() {
	var j = 10
	var i = 10
	fmt.Println(i, j) // there is space in between by default

	var x string = "hello"
	y := "by"

	fmt.Print(x, y) // observe there is no space in between
	fmt.Println()
	// so the variables are intgers print automatically addds space but when it is string we have to add it by our selvs
	fmt.Print(x, " ", y)
	fmt.Println()
	//but when we use println no need to add manually
	fmt.Println(x, y)
}

func verb() {
	yz := "hello"
	zy := 22
	fmt.Printf("The value of the zy variable is %v and the Type is %T\n", zy, zy)
	fmt.Printf("The value of the yz variable is %v and the Type is %T\n", yz, yz)
	fmt.Printf("%10d\n", zy) //adds 10 spacing from right
	fmt.Printf("%04d\n", zy) // checks number is 4 digit or not if not adds remaing zeros to it
}
