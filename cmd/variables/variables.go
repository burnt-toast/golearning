package main

import "fmt"
import "math"

// Cant use short hand notation on constants :=
//can let go infer the type
const s string = "constant"

func main() {

	//print constant value
	fmt.Println(s)

	const n = 500000000
	fmt.Println(n)

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

	//init a string variable
	var a string = "initial"
	fmt.Println(a)

	//init multiple int variables on one line
	var b, c int = 1, 2
	fmt.Println(b + c)

	//let go infer the type of the variable
	var f = true
	fmt.Println(f)

	//because e is not initialized it will get the 0 value which is a 0 for in
	var e int
	fmt.Println(e)

	//short hand for declaring and initializing
	z := "short"
	fmt.Println(z)

}
