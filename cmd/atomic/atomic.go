package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	// kick off three routines that will increase the count by 1 every second
	for i := 0; i < 3; i++ {
		fmt.Println("entering loop", i)
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				fmt.Println("in go routine")

				time.Sleep(time.Second)
			}
		}()
		fmt.Println("leaving loop", i)
	}
	fmt.Println("outside loop")
	time.Sleep(time.Second * 5)
	// sleeping for 5 seconds so count should be ~15 with the 3 workers
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("Final ops", opsFinal)
}
