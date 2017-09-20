package main

import (
	"strings"
	"fmt"
)

func main()  {
	data := []string{"a/b/c/d.go", "etc/home.go"}
	
	for _, val := range data {
		s := basename_2(val)
		fmt.Printf("\"%s\" string base is: %s", val, s)
		fmt.Println()
		
	}
}

func basename_2(str string) string {
	
	slash := strings.LastIndex(str, "/") // -1 if "/" not found.
	str  = str[slash+1:]
	
	if dot := strings.LastIndex(str, "."); dot >= 0 {
		str = str[:dot]
	}
	
	return str
}