package main

import "fmt"

func main() {

	//using the make func create a slice of strings with a length of 3
	s := make([]string, 3)
	fmt.Println("empty: ", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[1])

	fmt.Println(len(s))

	//append data to slice
	s = append(s, "d")
	s = append(s, "e")

	fmt.Println("append: ", s)

	//copy slice s -> c
	c := make([]string, 3)
	copy(c, s)
	fmt.Println("C: ", c)

	//playing with slice function
	fmt.Println("reminder of s: ", s)
	l := s[2:5]
	fmt.Println("s2:5", l)

	l = s[:5]
	fmt.Println("s:5", l)

	l = s[2:]
	fmt.Println("s2:", l)

	//declare and initializing
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	//multid
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
