package main

import (
	"os"
	"fmt"
)

//this program prints its command line arguments in each line.

func main()  {
	var s, sep string

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
