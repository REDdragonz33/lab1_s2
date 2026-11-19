package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	head *Node
}

func NewStack() *Stack { return &Stack{} }

func (s *Stack) GetType() byte { return 'S' }

func (s *Stack) Execute(cmd string, args []string) {
	switch cmd {
	case "SPUSH":
		if len(args) < 1 { fmt.Println("SPUSH: missing value"); return }
		s.head = NewNode(args[0], s.head)
	case "SPOP":
		if s.head == nil { fmt.Println("EMPTY"); return }
		fmt.Println(s.head.Data)
		s.head = s.head.Next
	}
}

func (s *Stack) Print() {
	fmt.Print("Stack (top -> bottom): ")
	cur := s.head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Println()
}

func (s *Stack) Save() string {
	var vals []string
	cur := s.head
	for cur != nil {
		vals = append(vals, cur.Data)
		cur = cur.Next
	}
	return fmt.Sprintf("%d %s", len(vals), strings.Join(vals, " "))
}

func (s *Stack) Load(data string) {
	s.head = nil
	parts := strings.Fields(data)
	if len(parts) == 0 { return }
	var count int
	fmt.Sscanf(parts[0], "%d", &count)
	for i := 1; i <= count && i < len(parts); i++ {
		s.head = NewNode(parts[i], s.head)
	}
}