package main

import (
	"fmt"
	"strings"
)

type DNode struct {
	Data string
	Prev *DNode
	Next *DNode
}

type DList struct {
	head *DNode
	tail *DNode
}

func NewDList() *DList { return &DList{} }

func (l *DList) GetType() byte { return 'L' }

func (l *DList) insertHead(val string) {
	n := &DNode{Data: val, Next: l.head}
	if l.head != nil { l.head.Prev = n } else { l.tail = n }
	l.head = n
}

func (l *DList) insertTail(val string) {
	n := &DNode{Data: val, Prev: l.tail}
	if l.tail != nil { l.tail.Next = n } else { l.head = n }
	l.tail = n
}

func (l *DList) Execute(cmd string, args []string) {
	switch cmd {
	case "LPUSH":
		if len(args) < 2 { fmt.Println("LPUSH: missing args"); return }
		mode, val := args[0], args[1]
		switch mode {
		case "HEAD": l.insertHead(val)
		case "TAIL": l.insertTail(val)
		case "AFTER":
			if len(args) < 3 { return }
			cur := l.head
			for cur != nil && cur.Data != args[2] { cur = cur.Next }
			if cur != nil {
				n := &DNode{Data: val, Prev: cur, Next: cur.Next}
				if cur.Next != nil { cur.Next.Prev = n } else { l.tail = n }
				cur.Next = n
			}
		case "BEFORE":
			if len(args) < 3 { return }
			cur := l.head
			for cur != nil && cur.Data != args[2] { cur = cur.Next }
			if cur != nil {
				n := &DNode{Data: val, Prev: cur.Prev, Next: cur}
				if cur.Prev != nil { cur.Prev.Next = n } else { l.head = n }
				cur.Prev = n
			}
		}
	case "LDEL":
		if len(args) < 1 { return }
		
		// Поддержка новых режимов: LDEL имя AFTER значение
		if len(args) >= 2 {
			mode := args[0]
			val := args[1]
			if mode == "AFTER" {
				// Удаление после
				cur := l.head
				for cur != nil && cur.Data != val { cur = cur.Next }
				if cur != nil && cur.Next != nil {
					temp := cur.Next
					cur.Next = temp.Next
					if temp.Next != nil { temp.Next.Prev = cur } else { l.tail = cur }
				}
				return
			} else if mode == "BEFORE" {
				// Удаление перед
				cur := l.head
				for cur != nil && cur.Data != val { cur = cur.Next }
				if cur != nil && cur.Prev != nil {
					temp := cur.Prev
					cur.Prev = temp.Prev
					if temp.Prev != nil { temp.Prev.Next = cur } else { l.head = cur }
				}
				return
			}
		}

		// Обычное удаление по значению
		val := args[0]
		cur := l.head
		for cur != nil && cur.Data != val { cur = cur.Next }
		if cur != nil {
			if cur.Prev != nil { cur.Prev.Next = cur.Next } else { l.head = cur.Next }
			if cur.Next != nil { cur.Next.Prev = cur.Prev } else { l.tail = cur.Prev }
		}
	}
}

func (l *DList) Print() {
	fmt.Print("DList: ")
	cur := l.head
	for cur != nil {
		fmt.Print(cur.Data, " <-> ")
		cur = cur.Next
	}
	fmt.Println("nullptr")
}

func (l *DList) Save() string {
	var vals []string
	cur := l.head
	for cur != nil {
		vals = append(vals, cur.Data)
		cur = cur.Next
	}
	return fmt.Sprintf("%d %s", len(vals), strings.Join(vals, " "))
}

func (l *DList) Load(data string) {
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