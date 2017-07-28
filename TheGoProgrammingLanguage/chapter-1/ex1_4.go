package main

/**
Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
*/

import (
	"os"
	"bufio"
	"fmt"
)

func main() {
	files := os.Args[1:]

	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup-2: %v\n", err)
			continue
		}
		var flag bool = checkDuplicateLines(f)
		if flag {
			fmt.Println(arg)
		}
		f.Close()
	}
	//for line, n := range counts {
	//	//fmt.Println("am here")
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
}

func checkDuplicateLines(f *os.File) bool {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if counts[line] > 1 {
			return true
		}

	}
	return false
}
