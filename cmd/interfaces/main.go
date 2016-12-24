package main

import (
	"github.com/burnt-toast/golearning/model"
)

func main() {

	r := model.Rect{Width: 3, Height: 4}
	c := model.Circle{Radius: 5}
	model.Measure(r)
	model.Measure(c)
}
