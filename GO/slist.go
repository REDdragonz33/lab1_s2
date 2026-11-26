package main

import (
	"errors"
	"fmt"
)

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

// --- ДОБАВЛЕНИЕ (4 способа) ---
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

// --- УДАЛЕНИЕ (4 способа, как в ТЗ) ---
func (l *SList) RemoveHead() {
	if l.Head != nil {
		l.Head = l.Head.Next
		if l.Head == nil {
			l.Tail = nil
		}
	}
}

func (l *SList) RemoveTail() {
	if l.Head == nil {
		return
	}
	if l.Head == l.Tail {
		l.Head = nil
		l.Tail = nil
		return
	}
	curr := l.Head
	for curr.Next != l.Tail {
		curr = curr.Next
	}
	curr.Next = nil
	l.Tail = curr
}

func (l *SList) RemoveAfter(target string) {
	curr := l.Head
	for curr != nil && curr.Value != target {
		curr = curr.Next
	}
	if curr != nil && curr.Next != nil {
		curr.Next = curr.Next.Next
		if curr.Next == nil {
			l.Tail = curr
		}
	}
}

func (l *SList) RemoveBefore(target string) {
	if l.Head == nil || l.Head.Value == target {
		return
	}
	if l.Head.Next != nil && l.Head.Next.Value == target {
		l.RemoveHead()
		return
	}
	curr := l.Head
	for curr.Next != nil && curr.Next.Next != nil && curr.Next.Next.Value != target {
		curr = curr.Next
	}
	if curr.Next != nil && curr.Next.Next != nil {
		curr.Next = curr.Next.Next
	}
}

// --- ЧТЕНИЕ И ПОИСК ---
func (l *SList) Search(val string) bool {
	curr := l.Head
	for curr != nil {
		if curr.Value == val {
			return true
		}
		curr = curr.Next
	}
	return false
}

func (l *SList) GetHead() (string, error) {
	if l.Head == nil {
		return "", errors.New("list is empty")
	}
	return l.Head.Value, nil
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