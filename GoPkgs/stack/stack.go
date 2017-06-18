package stack

/*
Stack as package

See the Stack exercise. In this exercise we want to create a separate package for that code.
Create a proper package for your stack implementation, Push, Pop and the Stack type need to be exported.

Write a simple unit test for this package. You should at least test that a Pop works after a Push.
*/

type Stack struct {
	index int
	values [10]int
}

//Push pushes an item on top of the stack.
func (s *Stack) Push(k int)  {
	s.values[s.index] = k
	s.index++
}

// Pop pops an item from the stack.
func (s *Stack) Pop() (ret int)  {
	s.index--
	ret = s.values[s.index]
	return
}