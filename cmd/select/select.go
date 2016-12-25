package main

import "time"
import "fmt"

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	c1Hit := false
	c2Hit := false

	for {
		fmt.Println("Iteration in for loop")
		if c1Hit && c2Hit {
			return
		}
		select {
		case m1 := <-c1:
			fmt.Println("Recieved", m1)
			c1Hit = true
		case m2 := <-c2:
			fmt.Println("Recieved", m2)
			c2Hit = true
		}
	}

}
