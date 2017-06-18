package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

//Create a reverse polish calculator. Uses stack package.

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)

type Stack struct {
	index int
	values [10]int
}

func (s *Stack) Push(k int)  {
	if s.index > 9 {
		return
	}
	s.values[s.index] = k
	s.index++
}

func (s *Stack) pop() (ret int) {
	s.index--
	if s.index < 0 {
		s.index = 0
		return
	}
	ret = s.values[s.index]
	return
}

func main()  {
	for {
		s, err := reader.ReadString('\n')
		var token string
		if err != nil {
			return
		}

		for _,c := range s {
			switch {
			case c >= '0' && c <= '9':
				token = token + string(c)
			case c == ' ':
				r, _ := strconv.Atoi(token)
				st.Push(r)
				token = ""
			case c == '+':
				fmt.Printf("%d\n", st.pop()+st.pop())
			case c == '-':
				p := st.pop()
				q := st.pop()
				fmt.Printf("%d\n", q - p)
			case c == '*':
				fmt.Printf("%d\n", st.pop() * st.pop())
			case c == 'q':
				return
			default:
				//error
			}
		}

	}
}

