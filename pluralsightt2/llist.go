package main

import (
	"errors"
	"fmt"
)

func main() {
	list := LinkedList{}
	list.Add(3)
	list.Add(5)
	list.Add(9)
	list.Add(2)

	fmt.Println()

	item, err := list.Pop()
	fmt.Println("Item:", item) // 9

	item, err = list.Pop()
	fmt.Println("Item:", item) // 5

	item, err = list.Pop()
	fmt.Println("Item:", item) // 3

	item, err = list.Pop()
	fmt.Println("Item:", item) // 3

	item, err = list.Pop()
	if err != nil {
		fmt.Println("An error ocurred retrieving next item from list:", err)
	}
}

// LinkedList is our implementation
type LinkedList struct {
	head *Node
}

// Add a new node to the top of the list (stack "flavor")
func (l *LinkedList) Add(data int) {
	fmt.Println("\nAdd - data to add is", data)
	if l.head == nil {
		l.head = &Node{data, nil} // this is like C#'s `head = new Node(){Data: data, Next: null}` -- it creates an "instance"
		fmt.Printf("Add - head address (&l.head) is %p\n", &l.head)
		fmt.Printf("Add - head is &{Data: %v, Next: %v}\n", l.head.Data, l.head.Next)

	} else {
		fmt.Printf("Add - head is &{Data: %d, Next: %p}\n", l.head.Data, l.head.Next)
		fmt.Printf("Add - head is &{Data: %v, Next: %v}\n", l.head.Data, l.head.Next)

		fmt.Println("\tnode := &Node{data, nil}")

		node := &Node{data, nil}
		fmt.Printf("Add - node is &{Data: %d, Next: %p}\n", node.Data, node.Next)
		fmt.Printf("Add - node is &{Data: %v, Next: %v}\n", node.Data, node.Next)

		fmt.Println("\tnode.Next = l.head")

		node.Next = l.head

		fmt.Printf("Add - node is &{Data: %d, Next: %p}\n", node.Data, node.Next)
		fmt.Printf("Add - node is &{Data: %v, Next: %v}\n", node.Data, node.Next)

		fmt.Println("\tl.head = node")
		l.head = node
		fmt.Printf("Add - head is &{Data: %d, Next: %p}\n", l.head.Data, l.head.Next)
		fmt.Printf("Add - head is &{Data: %v, Next: %v}\n", l.head.Data, l.head.Next)
	}
}

// Pop returns the item at the top/head
func (l *LinkedList) Pop() (int, error) {
	if l.head == nil {
		return 0, errors.New("Head cannot be null")
	}
	item := l.head.Data
	l.head = l.head.Next
	return item, nil
}

// Node helps build the linked list
type Node struct {
	Data int
	Next *Node
}
