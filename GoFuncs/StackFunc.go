package main

import (
	"fmt"
	"strconv"
)

/*
Stack

Create a simple stack which can hold a fixed number of ints. It does not have to grow beyond this limit.
Define push -- put something on the stack -- and pop -- retrieve something from the stack -- functions.
The stack should be a LIFO (last in, first out) stack.

Write a String method which converts the stack to a string representation. The stack in the figure could be
represented as: [0:m] [1:l] [2:k] .
 */

type stack struct {
	index int
	data [10] int
}

func (s *stack) push(k int)  {
	s.data[s.index] = k
	s.index++
}

func (s *stack) pop() int  {
	s.index--;
	ret := s.data[s.index]
	s.data[s.index] = 0
	return ret
}

func (s stack) String() string  {
	var str string
	for i := 0; i<=s.index ; i++  {
		str = str + "[" + strconv.Itoa(i) + ":" + strconv.Itoa(s.data[i])+ "]"
	}
	return str

}

func main()  {
	var s stack
	s.push(40)
	s.push(50)
	fmt.Printf("stack: %v\n", s)
	res := s.String()
	fmt.Println(res)
}
