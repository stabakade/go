package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
		close(c1) // very important to close
	}()

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- 2
		close(c2) // always remember to close channel when using
	}()

// standard tenplate for for-select
	for {
    // remember to add this check before select
		if c1 == nil && c2 == nil {
			break
		}

		select {
		case val1, ok := <-c1:
			if ok {
				fmt.Println(val1)
			} else {
				c1 = nil // very important to set to nil yourself else deadlock
			}
		case val2, ok := <-c2:
			if ok {
				fmt.Println(val2)
			} else {
				c2 = nil
			}
		}
	}

	fmt.Println("execution done")
}
