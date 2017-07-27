/*
Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
 */
package main

import (
	"os"
	"fmt"
)

func main()  {

	for index, value := range os.Args[0:] {
		fmt.Println(index, value)
	}
}
