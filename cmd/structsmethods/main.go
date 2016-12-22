package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

// calculate area of rectangle via the pointer
func (r *rect) area() int {
	return r.width * r.height
}

// calculate permiter with a value reciever
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (rect) print() string {
	return "You are using the rect struct"
}

func main() {
	r := rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim", r.perim())

	rp := &r

	fmt.Println("area:", rp.area())
	fmt.Println("perim", rp.perim())

	fmt.Println(r.print())
}
