package main

import (
	"errors"
	"fmt"
)

type QueueNode struct {
	Value string
	Next  *QueueNode
}

type Queue struct {
	Head *QueueNode
	Tail *QueueNode
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Push(val string) {
	newNode := &QueueNode{Value: val}
	if q.Tail != nil {
		q.Tail.Next = newNode
	} else {
		q.Head = newNode
	}
	q.Tail = newNode
}

func (q *Queue) Pop() (string, error) {
	if q.Head == nil {
		return "", errors.New("queue is empty")
	}
	val := q.Head.Value
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil // Очередь опустела
	}
	return val, nil
}

func (q *Queue) Print() {
	fmt.Printf("Queue (front): ")
	curr := q.Head
	for curr != nil {
		fmt.Printf("%s ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}