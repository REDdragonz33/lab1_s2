package main

import (
	"errors"
	"fmt"
)

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

// --- ДОБАВЛЕНИЕ (4 способа) ---
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

// --- УДАЛЕНИЕ (4 способа, как в ТЗ) ---
func (l *DList) RemoveHead() {
	if l.Head != nil {
		l.Head = l.Head.Next
		if l.Head != nil {
			l.Head.Prev = nil
		} else {
			l.Tail = nil
		}
	}
}

func (l *DList) RemoveTail() {
	if l.Tail != nil {
		l.Tail = l.Tail.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		} else {
			l.Head = nil
		}
	}
}

func (l *DList) RemoveAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Value != target {
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		curr.Next = curr.Next.Next
		if curr.Next != nil {
			curr.Next.Prev = curr
		} else {
			l.Tail = curr
		}
	}
}

func (l *DList) RemoveBefore(target string) {
	curr := l.Head
	for curr != nil && curr.Value != target {
		curr = curr.Next
	}
	if curr != nil && curr.Prev != nil {
		prev := curr.Prev
		if prev.Prev != nil {
			prev.Prev.Next = curr
			curr.Prev = prev.Prev
		} else {
			l.Head = curr
			curr.Prev = nil
		}
	}
}

// --- ЧТЕНИЕ И ПОИСК ---
func (l *DList) Search(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Value == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *DList) GetHead() (string, error) {
	if l.Head == nil {
		return "", errors.New("list is empty")
	}
	return l.Head.Value, nil
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