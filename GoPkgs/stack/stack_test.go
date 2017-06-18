package stack

import "testing"

func TestPushPop(t *testing.T)  {
	stack := new(Stack)

	stack.Push(5)
	if stack.Pop() != 5 {
		t.Log("Pop doesnt give 5")
		t.Fail()
	}
}
