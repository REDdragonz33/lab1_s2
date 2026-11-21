package main

import "fmt"

type SNode struct {
	Value string
	Next  *SNode
}

type SList struct {
	Head *SNode
	Tail *SNode
}

func NewSList() *SList {
	return &SList{}
}

func (l *SList) InsertHead(val string) {
	newNode := &SNode{Value: val, Next: l.Head}
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}

func (l *SList) InsertTail(val string) {
	newNode := &SNode{Value: val}
	if l.Tail != nil {
		l.Tail.Next = newNode
	} else {
		l.Head = newNode
	}
	l.Tail = newNode
}

func (l *SList) InsertAfter(target, val string) {
	curr := l.Head
	for curr != nil && curr.Value != target {
		curr = curr.Next
	}
	if curr != nil {
		newNode := &SNode{Value: val, Next: curr.Next}
		curr.Next = newNode
		if curr == l.Tail {
			l.Tail = newNode
		}
	}
}

func (l *SList) InsertBefore(target, val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == target {
		l.InsertHead(val)
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Value != target {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = &SNode{Value: val, Next: curr.Next}
	}
}

func (l *SList) Remove(val string) {
	if l.Head == nil {
		return
	}
	if l.Head.Value == val {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Value != val {
		curr = curr.Next
	}
	if curr.Next != nil {
		curr.Next = curr.Next.Next
		if curr.Next == nil {
			l.Tail = curr
		}
	}
}

func (l *SList) Print() {
	fmt.Printf("SList: ")
	curr := l.Head
	for curr != nil {
		fmt.Printf("%s -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}