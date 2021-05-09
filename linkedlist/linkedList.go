package main

import (
	"fmt"
)

type Node struct {
	value int
	next *Node
}

type Stack struct {
	head *Node
}

func NewStack() *Stack {
	return &Stack{
		head : nil,
	}
}

func NewNode(n int) *Node {
	return &Node{
		value : n,
		next : nil,
	}
}

func (l *Stack) Push(n int)  {
	newNode := NewNode(n)
	if l.head == nil {
		l.head = newNode
		return
	}
	newNode.next = l.head
	l.head = newNode
}

func (l Stack) Display() {
	for {
		if l.head == nil {
			fmt.Println("Empty Stack")
			break
		} else {
			fmt.Printf("%d", l.head.value)
			if l.head.next != nil {
				fmt.Printf(" -> ")
				l.head = l.head.next
			} else {
				fmt.Printf("\n")
				break
			}
		}
	}
}

func (l *Stack) Pop()  interface{} {
	var currentNode *Node
	if l.head == nil {
		fmt.Println("Error: empty stack")
		return nil
	}
	currentNode = l.head
	l.head = l.head.next
	currentNode.next = nil
	return currentNode.value
}

func main() {
	stack := NewStack()
	stack.Push(50)
	stack.Push(100)
	stack.Push(150)
	stack.Push(200)
	stack.Push(250)
	stack.Display()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	stack.Display()
}