package main

import "fmt"

func main()  {
	n := []int{5, -1, 0, 12, 3, 7}
	fmt.Printf("unsorted: %v\n", n)
	bubblesort(n)
	fmt.Printf("sorted: %v", n)

}

func bubblesort(n [] int) {
	for i := 0; i < len(n)-1; i++ {
		for j := i+1; j < len(n); j++ {
			if n[j] < n[i] {
				n[j], n[i] = n[i], n[j]
			}
		}
	}
}
