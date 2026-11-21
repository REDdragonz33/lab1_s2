package main

import (
	"errors"
	"fmt"
)

type StackNode struct {
	Value string
	Next  *StackNode
}

type Stack struct {
	Top *StackNode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val string) {
	s.Top = &StackNode{Value: val, Next: s.Top}
}

func (s *Stack) Pop() (string, error) {
	if s.Top == nil {
		return "", errors.New("stack is empty")
	}
	val := s.Top.Value
	s.Top = s.Top.Next
	return val, nil
}

func (s *Stack) Print() {
	fmt.Printf("Stack (top): ")
	curr := s.Top
	for curr != nil {
		fmt.Printf("%s ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}