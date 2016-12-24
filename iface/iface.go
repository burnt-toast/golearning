package iface

import "fmt"

type Geomtry interface {
	Area() float64
	Perim() float64
}

func Measure(g Geomtry) {
	fmt.Println(g)
	fmt.Println(g.Area())
	fmt.Println(g.Perim())
}
