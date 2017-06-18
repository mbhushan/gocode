package main

import (
	"container/list"
	"fmt"
)

/*
Make use of the package container/list to create a (doubly) linked list.
Push the values 1, 2 and 4 to the list and then print it.
*/

func main()  {
	dl := list.New()
	dl.PushBack(1)
	dl.PushBack(2)
	dl.PushBack(4)

	for node := dl.Front(); node != nil; node = node.Next()  {
		fmt.Printf("%v\n", node.Value)
	}
}
