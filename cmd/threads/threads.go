package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")
	fmt.Println("test")
	go f("routine 2")
	fmt.Println("test2")
	go func(msg string) {
		fmt.Println(msg)
	}("anonymous")
	fmt.Println("test3")
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
