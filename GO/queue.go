package main

import (
	"fmt"
	"strings"
)

type Queue struct {
	head *Node
	tail *Node
}

func NewQueue() *Queue { return &Queue{} }

func (q *Queue) GetType() byte { return 'Q' }

func (q *Queue) Execute(cmd string, args []string) {
	switch cmd {
	case "QPUSH":
		if len(args) < 1 { fmt.Println("QPUSH: missing value"); return }
		n := NewNode(args[0], nil)
		if q.tail != nil {
			q.tail.Next = n
		} else {
			q.head = n
		}
		q.tail = n
	case "QPOP":
		if q.head == nil { fmt.Println("EMPTY"); return }
		fmt.Println(q.head.Data)
		q.head = q.head.Next
		if q.head == nil { q.tail = nil }
	}
}

func (q *Queue) Print() {
	fmt.Print("Queue (front -> back): ")
	cur := q.head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func (q *Queue) Save() string {
	var vals []string
	cur := q.head
	for cur != nil {
		vals = append(vals, cur.Data)
		cur = cur.Next
	}
	return fmt.Sprintf("%d %s", len(vals), strings.Join(vals, " "))
}

func (q *Queue) Load(data string) {
	q.head = nil
	q.tail = nil
	parts := strings.Fields(data)
	if len(parts) == 0 { return }
	var count int
	fmt.Sscanf(parts[0], "%d", &count)
	for i := 1; i <= count && i < len(parts); i++ {
		q.Execute("QPUSH", []string{parts[i]})
	}
}