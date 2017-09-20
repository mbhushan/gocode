package main

import "fmt"

//comma inserts comma in a non-negative decimal integer string.

func main()  {

	data := []string{"12345", "234567", "12345678", "1234567890"}
	
	for _, val := range data {
		fmt.Println(val," after 3 comma: ", comma(val))
	}
}

func comma(str string) string {
	n := len(str)
	if n <= 3 {
		return str
	}
	
	return comma(str[:n-3]) + "," + str[n-3:]
}
