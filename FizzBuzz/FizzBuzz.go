package main

import "fmt"

func main() {

	const (
		FIZZ = 3
		BUZZ = 5
	)

	for i := 1; i<100; i++  {
		if i%FIZZ ==0 && i%BUZZ == 0 {
			fmt.Printf("FizzBuzz")
		} else if i%FIZZ == 0 {
			fmt.Printf("Fizz")
		} else if i%BUZZ == 0 {
			fmt.Printf("Buzz")
		} else {
			fmt.Printf("%d", i)
		}
		fmt.Println()
	}

}
