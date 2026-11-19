package main

import "fmt"

type IStructure interface {
	Execute(cmd string, args []string)
	Print()
	Save() string
	Load(data string)
	GetType() byte
}

type Node struct {
	Data string
	Next *Node
}

func NewNode(data string, next *Node) *Node {
	return &Node{Data: data, Next: next}
}

func ParseInt(s string) (int, bool) {
	var res int
	_, err := fmt.Sscanf(s, "%d", &res)
	return res, err == nil
}