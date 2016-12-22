package main

import "fmt"

func main() {

	i := 1
	//loop with single condition
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//initial/condition/after loop
	for j := 7; j < 9; j++ {
		fmt.Println(j)
	}

	//for until broken

	for {
		fmt.Println("loop")
		break
	}

}
