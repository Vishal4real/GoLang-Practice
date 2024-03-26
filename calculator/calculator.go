package main

import (
	"fmt"
)

func main() {
	var choice, num1, num2 int
	fmt.Println("This is Simple calculator \n Enter Your choice number")
	fmt.Println("1. Addition \n2. Substraction \n3. Multiplication \n4. Division")
	fmt.Scanln(&choice)
	if choice > 4 || choice < 1 {
		fmt.Print("Please select proper input number")
		main()
	} else {
		takeip(&num1, &num2)
		calculate(choice, num1, num2)
	}
}
func takeip(num1, num2 *int) {
	fmt.Println("Enter Number 1 & 2")
	fmt.Scan(num1, num2)
}
func calculate(c, n1, n2 int) {
	switch {
	case c == 1:
		fmt.Printf("Addition of %v and %v is %v", n1, n2, n1+n2)
	case c == 2:
		fmt.Printf("Substraction of %v and %v is %v", n1, n2, n1-n2)
	case c == 3:
		fmt.Printf("Multiplication of %v and %v is %v", n1, n2, n1*n2)
	case c == 4:
		fmt.Printf("Divison of %v and %v is %v", n1, n2, n1/n2)
	default:
		fmt.Printf("Select Proper Option")
	}
}
