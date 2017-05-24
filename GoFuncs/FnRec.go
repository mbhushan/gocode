package main

import "fmt"

func main ()  {
	var i int = 0
	rec(i)
}

func rec(i int)  {
	if i == 10 {
		return
	}
	rec(i+1)
	fmt.Printf("%d ", i)
}
