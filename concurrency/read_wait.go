package main

import (
	"time"
	"fmt"
)

func ready(w string, sec int)  {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I am waiting!")
	time.Sleep(5*time.Second)
}

