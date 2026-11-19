package main

import (
	"fmt"
	"strings"
)

type SList struct {
	head *Node
	tail *Node
}

func NewSList() *SList { return &SList{} }

func (l *SList) GetType() byte { return 'F' }

func (l *SList) insertHead(val string) {
	n := NewNode(val, l.head)
	l.head = n
	if l.tail == nil { l.tail = n }
}

func (l *SList) insertTail(val string) {
	n := NewNode(val, nil)
	if l.tail != nil {
		l.tail.Next = n
	} else {
		l.head = n
	}
	l.tail = n
}

func (l *SList) Execute(cmd string, args []string) {
	switch cmd {
	case "FPUSH":
		if len(args) < 2 { fmt.Println("FPUSH: missing args"); return }
		mode, val := args[0], args[1]
		switch mode {
		case "HEAD":
			l.insertHead(val)
		case "TAIL":
			l.insertTail(val)
		case "AFTER":
			if len(args) < 3 { return }
			cur := l.head
			for cur != nil && cur.Data != args[2] { cur = cur.Next }
			if cur != nil {
				n := NewNode(val, cur.Next)
				cur.Next = n
				if cur == l.tail { l.tail = n }
			}
		case "BEFORE":
			if len(args) < 3 { return }
			if l.head != nil && l.head.Data == args[2] { l.insertHead(val); return }
			cur := l.head
			for cur != nil && cur.Next != nil && cur.Next.Data != args[2] { cur = cur.Next }
			if cur != nil && cur.Next != nil {
				cur.Next = NewNode(val, cur.Next)
			}
		}
	case "FDEL":
		if len(args) < 1 { return }
		
		// Поддержка новых режимов: FDEL имя AFTER значение
		if len(args) >= 2 {
			mode := args[0]
			val := args[1]
			if mode == "AFTER" {
				// Удаление после
				cur := l.head
				for cur != nil && cur.Data != val { cur = cur.Next }
				if cur != nil && cur.Next != nil {
					if cur.Next == l.tail { l.tail = cur }
					cur.Next = cur.Next.Next
				}
				return
			} else if mode == "BEFORE" {
				// Удаление перед
				if l.head == nil || l.head.Next == nil { return }
				if l.head.Next.Data == val {
					l.head = l.head.Next
					return
				}
				cur := l.head
				for cur.Next != nil && cur.Next.Next != nil {
					if cur.Next.Next.Data == val {
						cur.Next = cur.Next.Next
						return
					}
					cur = cur.Next
				}
				return
			}
		}

		// Обычное удаление по значению
		val := args[0]
		if l.head == nil { return }
		if l.head.Data == val {
			l.head = l.head.Next
			if l.head == nil { l.tail = nil }
			return
		}
		cur := l.head
		for cur.Next != nil && cur.Next.Data != val { cur = cur.Next }
		if cur.Next != nil {
			if cur.Next == l.tail { l.tail = cur }
			cur.Next = cur.Next.Next
		}
		
	case "FGET":
		// Поиск по значению (выводит TRUE/FALSE)
		if len(args) < 1 { return }
		val := args[0]
		cur := l.head
		for cur != nil {
			if cur.Data == val {
				fmt.Println("TRUE")
				return
			}
			cur = cur.Next
		}
		fmt.Println("FALSE")
	}
}

func (l *SList) Print() {
	fmt.Print("SList: ")
	cur := l.head
	for cur != nil {
		fmt.Print(cur.Data, " -> ")
		cur = cur.Next
	}
	fmt.Println("nullptr")
}

func (l *SList) Save() string {
	var vals []string
	cur := l.head
	for cur != nil {
		vals = append(vals, cur.Data)
		cur = cur.Next
	}
	return fmt.Sprintf("%d %s", len(vals), strings.Join(vals, " "))
}

func (l *SList) Load(data string) {
	l.head = nil
	l.tail = nil
	parts := strings.Fields(data)
	if len(parts) == 0 { return }
	var count int
	fmt.Sscanf(parts[0], "%d", &count)
	for i := 1; i <= count && i < len(parts); i++ {
		l.insertTail(parts[i])
	}
}