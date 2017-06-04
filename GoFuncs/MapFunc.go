package main

import "fmt"

/*
Map function

A map()-function is a function that takes a function and a list. The function is applied to each member in the list
and a new list containing these calculated values is returned. Thus:

map(f(),(a1,a2,…,an−1,an))=(f(a1),f(a2),…,f(an−1),f(an))
map(f(),(a1,a2,…,an−1,an))=(f(a1),f(a2),…,f(an−1),f(an))

Write a simple map()-function in Go. It is sufficient for this function only to work for ints.
*/

func doMap(f func(int) int, A [] int)  [] int{
	B := make([]int, len(A))
	for k, v := range A {
		B[k] = f(v)
	}
	return B
}

func main()  {
	A := []int{1,2,3,4,5,6}
	f := func(i int) int {
		return i*i
	}
	fmt.Printf("%v\n", doMap(f, A))
}
