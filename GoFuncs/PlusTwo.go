package main

import "fmt"

func main()  {
	p2 := plusTwo()
	x := []int{1,2,3,4,5}
	for i := 0; i<len(x); i++  {
		fmt.Printf("%d\n", p2(x[i]))
	}
}

func plusTwo() func(int) int {
	return func (x int) int {return x + 2}
}