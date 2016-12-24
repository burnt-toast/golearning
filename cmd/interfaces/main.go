package main

import (
	"github.com/burnt-toast/golearning/iface"
	"github.com/burnt-toast/golearning/model"
)

func main() {

	r := model.Rect{Width: 3, Height: 4}
	c := model.Circle{Radius: 5}
	iface.Measure(r)
	iface.Measure(c)
}
