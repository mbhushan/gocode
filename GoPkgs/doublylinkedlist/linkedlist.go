package main

import (
	"errors"
	"fmt"
)

/*
Create your own linked list implementation. And perform the same actions as above.
*/

type Value int

type Node struct {
	Value
	prev, next *Node
}

type List struct {
	head, tail *Node
}

func (l *List) Front() *Node {
	return l.head
}

func (node *Node) Next() *Node  {
	return node.next
}

func (list *List) Push(val Value)  *List{
	node := &Node{Value: val}

	if list.head == nil {
		list.head = node
	} else {
		list.tail.next = node
		node.prev = list.tail
	}
	list.tail = node

	return list
}

var errEmpty = errors.New("List is empty")

func (list *List) Pop() (v Value, err error)  {
	if list.tail == nil {
		err = errEmpty
	} else {
		v = list.tail.Value
		list.tail = list.tail.prev
		if list.tail == nil {
			list.head = nil
		}
	}

	return v, err
}

func main()  {

	list := new(List)

	list.Push(1)
	list.Push(2)
	list.Push(4)

	for node := list.Front(); node != nil; node=node.next {
		fmt.Printf("%v\n", node.Value)
	}
	fmt.Println()

	for v, err := list.Pop(); err == nil; v, err = list.Pop() {
		fmt.Printf("%v\n", v)
	}

}
