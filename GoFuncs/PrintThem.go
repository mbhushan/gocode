package main

import "fmt"

func main()  {
	printThem(1, 2,3 , 4, 5)
	printThem(7,8,9,0)
}

func printThem(args ... int) {
	for _, n := range args {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
