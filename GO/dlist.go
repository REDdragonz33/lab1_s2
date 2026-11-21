package main

import "fmt"

type DNode struct {
	Value string
	Prev  *DNode
	Next  *DNode
}

type DList struct {
	Head *DNode
	Tail *DNode
}

func NewDList() *DList {
	return &DList{}
}

func (l *DList) InsertHead(val string) {
	newNode := &DNode{Value: val, Next: l.Head}
	if l.Head != nil {
		l.Head.Prev = newNode
	} else {
		l.Tail = newNode
	}
	l.Head = newNode
}

func (l *DList) InsertTail(val string) {
	newNode := &DNode{Value: val, Prev: l.Tail}
	if l.Tail != nil {
		l.Tail.Next = newNode
	} else {
		l.Head = newNode
	}
	l.Tail = newNode
}

func (l *DList) InsertAfter(target, val string) {
	curr := l.Head
	for curr != nil && curr.Value != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &DNode{Value: val, Prev: curr, Next: curr.Next}
		if curr.Next != nil {
			curr.Next.Prev = newNode
		} else {
			l.Tail = newNode
		}
		curr.Next = newNode
	}
}

func (l *DList) InsertBefore(target, val string) {
	curr := l.Head
	for curr != nil && curr.Value != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &DNode{Value: val, Prev: curr.Prev, Next: curr}
		if curr.Prev != nil {
			curr.Prev.Next = newNode
		} else {
			l.Head = newNode
		}
		curr.Prev = newNode
	}
}

func (l *DList) Remove(val string) {
	curr := l.Head
	for curr != nil && curr.Value != val {
		curr = curr.Next
	}
	if curr != nil {
		if curr.Prev != nil {
			curr.Prev.Next = curr.Next
		} else {
			l.Head = curr.Next
		}
		if curr.Next != nil {
			curr.Next.Prev = curr.Prev
		} else {
			l.Tail = curr.Prev
		}
	}
}

func (l *DList) Print() {
	fmt.Printf("DList: ")
	curr := l.Head
	for curr != nil {
		fmt.Printf("%s <-> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}