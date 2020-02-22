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

	fmt.Println()

	item, err := list.Pop()
	fmt.Println("Item:", item) // 9

	item, err = list.Pop()
	fmt.Println("Item:", item) // 5

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
func (l *LinkedList) Add(data int) { // attaching a method to the struct
	fmt.Println("\nAdd - data to add is", data)
	if l.head == nil {
		l.head = &Node{data, nil} // this is like C#'s `head = new Node(){Data: data, Next: null}` -- it creates an "instance"
		fmt.Printf("Add - head address (&l.head) is %p\n", &l.head)
		// fmt.Printf("Add - head is {Data: %d, Next: %p}\n", l.head.Data, l.head.Next)
	} else {
		// fmt.Printf("Add - head address (&l.head) is %p\n", &l.head)
		// fmt.Println("Add - l.head is", l.head)
		// fmt.Println("Add - l.head.Data is", l.zhead.Data)
		// fmt.Println("Add - l.head.Next is", l.head.Next)
		fmt.Printf("Add - head is {Data: %d, Next: %p}\n", l.head.Data, l.head.Next)

		node := &Node{data, nil}
		// fmt.Println("Add - node is", node)
		// fmt.Println("Add - node.Data is", node.Data)
		// fmt.Println("Add - node.Next is", node.Next)
		fmt.Printf("Add - node is {Data: %d, Next: %p}\n", node.Data, node.Next)

		fmt.Println("node.Next = l.head")

		node.Next = l.head

		fmt.Printf("Add - node is {Data: %d, Next: %p}\n", node.Data, node.Next)

		fmt.Println("l.head = node")
		l.head = node
		fmt.Printf("Add - head is {Data: %d, Next: %p}\n", l.head.Data, l.head.Next)
		fmt.Println("Add - head.Data is", l.head.Data)
		fmt.Println("Add - head.Next is", l.head.Next)

		// l.head = Node{Data: node.Data, Next: node.Next}
	}

	// if l.head == (Node{}) { // this is how we compre the head in "null"/empty (structs can't be nul in go)
	// 	l.head = Node{data, nil}
	// 	fmt.Println("If - head.Data is", l.head.Data)
	// 	fmt.Println("If - head.Next is nil", l.head.Next)
	// } else {
	// 	node := Node{data, nil}
	// 	node.Next = &l.head
	// 	l.head = node
	// 	fmt.Println("Else - head.Data is", l.head.Data)
	// 	fmt.Println("Else - head.Next is", l.head.Next)
	// }
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
