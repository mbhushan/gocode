package main

import "fmt"

type nameage struct {
	name string
	age int
}

func main()  {
	a := new(nameage)
	a.name = "Mani Bhushan"
	a.age = 21
	fmt.Printf("%v\n", a)
}
