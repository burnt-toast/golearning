package main

import "fmt"

func main() {

	//basic if else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("Seven is odd")
	}

	//statement can precede conditions, any variables declared are available to branches
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
