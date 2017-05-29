package main

import "fmt"

func main()  {
	A := []int {1,2,3,9,8,7,6,5,4}

	var ans int = maxInArr(A)
	fmt.Printf("max in A: %d", ans)
}

func maxInArr(A [] int) (max int) {
	max = A[0]

	for _,v := range A {
		if v > max {
			max = v
		}
	}
	return
}
