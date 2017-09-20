package main

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprint(values) but adds commas.

func main()  {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

func intsToString(values []int) string  {
	var buf bytes.Buffer
	buf.WriteByte('[')
	
	for index, val := range values {
		if index > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", val)
	}
	buf.WriteByte(']')
	return buf.String()
}


