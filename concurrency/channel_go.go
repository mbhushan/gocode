package main

import (
	"time"
	"fmt"
)

var c chan int

func ready(w string, sec int)  {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is about to start!")
	c <- 1
}

func main() {
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I am waiting, but not too long")
	i := 0
	L: for  {
		select {
		case <- c:
			i++
			if i > 1 {
				break L
			}
		}
		}
}
