package main

import (
	"even"
	"fmt"
)

func main() {
	i := 5
	fmt.Printf("is %d even? %v\n", i, even.Even(i))
}
