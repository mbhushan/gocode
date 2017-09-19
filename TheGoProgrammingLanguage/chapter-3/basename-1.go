package main

import "fmt"

func main() {
	data := []string{"a/b/c/d.go", "etc/home.go"}
	
	for _, val := range data {
		s := basename(val)
		fmt.Printf("\"%s\" string base is: %s", val, s)
		fmt.Println()
		
	}
}

func basename(s string)  string {
	
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1: ]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i>=0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
	
}
