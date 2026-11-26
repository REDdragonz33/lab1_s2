package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IStructure interface {
	Execute(cmd string, args []string)
	Print()
}

type Registry struct {
	structures map[string]IStructure
}

func NewRegistry() *Registry {
	return &Registry{
		structures: make(map[string]IStructure),
	}
}

func (r *Registry) GetOrCreate(name string, typeName byte) IStructure {
	if s, exists := r.structures[name]; exists {
		return s
	}
	
	var newStruct IStructure
	switch typeName {
	case 'M':
		newStruct = NewArray()
	case 'F':
		newStruct = NewSList()
	case 'L':
		newStruct = NewDList()
	case 'S':
		newStruct = NewStack()
	case 'Q':
		newStruct = NewQueue()
	// Дерево ('T') намеренно исключено, так как по ТЗ оно реализуется на C++
	default:
		return nil
	}
	
	r.structures[name] = newStruct
	return newStruct
}

// --- МАССИВ (Array) ---
type Array struct {
	data []string
}

func NewArray() *Array {
	return &Array{data: make([]string, 0)}
}

func (a *Array) PushEnd(val string) {
	a.data = append(a.data, val)
}

func (a *Array) PushAt(idx int, val string) {
	if idx < 0 || idx > len(a.data) {
		fmt.Println("Index out of range")
		return
	}
	a.data = append(a.data, "")
	copy(a.data[idx+1:], a.data[idx:])
	a.data[idx] = val
}

func (a *Array) RemoveAt(idx int) {
	if idx < 0 || idx >= len(a.data) {
		fmt.Println("Index out of range")
		return
	}
	a.data = append(a.data[:idx], a.data[idx+1:]...)
}

func (a *Array) Execute(cmd string, args []string) {
	switch cmd {
	case "MPUSH":
		if len(args) == 1 {
			a.PushEnd(args[0])
		} else if len(args) == 2 {
			idx, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("MPUSH: invalid index")
				return
			}
			a.PushAt(idx, args[1])
		}
	case "MDEL":
		if len(args) == 1 {
			idx, _ := strconv.Atoi(args[0])
			a.RemoveAt(idx)
		}
	case "MGET":
		if len(args) == 1 {
			idx, _ := strconv.Atoi(args[0])
			if idx >= 0 && idx < len(a.data) {
				fmt.Println(a.data[idx])
			} else {
				fmt.Println("Index out of range")
			}
		}
	case "MLEN":
		fmt.Println(len(a.data))
	default:
		fmt.Println("Unknown array command:", cmd)
	}
}

func (a *Array) Print() {
	fmt.Printf("Array[%d]: ", len(a.data))
	for _, v := range a.data {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}

// --- ОДНОСВЯЗНЫЙ СПИСОК (SList) ---
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

func (l *SList) Execute(cmd string, args []string) {
	switch cmd {
	case "FPUSH":
		if len(args) >= 2 {
			mode := args[0]
			val := args[1]
			if mode == "HEAD" {
				l.InsertHead(val)
			} else if mode == "TAIL" {
				l.InsertTail(val)
			} else if mode == "AFTER" && len(args) >= 3 {
				l.InsertAfter(val, args[2])
			} else if mode == "BEFORE" && len(args) >= 3 {
				l.InsertBefore(val, args[2])
			}
		}
	case "FDEL":
		if len(args) >= 1 {
			l.Remove(args[0])
		}
	default:
		fmt.Println("Unknown slist command:", cmd)
	}
}

func (l *SList) Print() {
	fmt.Print("SList: ")
	curr := l.Head
	for curr != nil {
		fmt.Printf("%s -> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// --- ДВУСВЯЗНЫЙ СПИСОК (DList) ---
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

func (l *DList) Remove(val string) {
	curr := l.Head
	for curr != nil && curr.Value != val {
		curr = curr.Next
	}
	if curr != nil {
		if curr.Prev != nil {
			curr.Prev.Next = curr.Next
		} else {
			l.Head = curr.Next
		}
		if curr.Next != nil {
			curr.Next.Prev = curr.Prev
		} else {
			l.Tail = curr.Prev
		}
	}
}

func (l *DList) Execute(cmd string, args []string) {
	switch cmd {
	case "LPUSH":
		if len(args) >= 2 {
			mode := args[0]
			val := args[1]
			if mode == "HEAD" {
				l.InsertHead(val)
			} else if mode == "TAIL" {
				l.InsertTail(val)
			} else if mode == "AFTER" && len(args) >= 3 {
				l.InsertAfter(val, args[2])
			} else if mode == "BEFORE" && len(args) >= 3 {
				l.InsertBefore(val, args[2])
			}
		}
	case "LDEL":
		if len(args) >= 1 {
			l.Remove(args[0])
		}
	default:
		fmt.Println("Unknown dlist command:", cmd)
	}
}

func (l *DList) Print() {
	fmt.Print("DList: ")
	curr := l.Head
	for curr != nil {
		fmt.Printf("%s <-> ", curr.Value)
		curr = curr.Next
	}
	fmt.Println("nil")
}

// --- СТЕК (Stack) ---
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

func (s *Stack) Pop() {
	if s.Top == nil {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println(s.Top.Value)
	s.Top = s.Top.Next
}

func (s *Stack) Execute(cmd string, args []string) {
	switch cmd {
	case "SPUSH":
		if len(args) >= 1 {
			s.Push(args[0])
		}
	case "SPOP":
		s.Pop()
	default:
		fmt.Println("Unknown stack command:", cmd)
	}
}

func (s *Stack) Print() {
	fmt.Print("Stack (top): ")
	curr := s.Top
	for curr != nil {
		fmt.Printf("%s ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}

// --- ОЧЕРЕДЬ (Queue) ---
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

func (q *Queue) Pop() {
	if q.Head == nil {
		fmt.Println("Queue is empty")
		return
	}
	fmt.Println(q.Head.Value)
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
}

func (q *Queue) Execute(cmd string, args []string) {
	switch cmd {
	case "QPUSH":
		if len(args) >= 1 {
			q.Push(args[0])
		}
	case "QPOP":
		q.Pop()
	default:
		fmt.Println("Unknown queue command:", cmd)
	}
}

func (q *Queue) Print() {
	fmt.Print("Queue (front): ")
	curr := q.Head
	for curr != nil {
		fmt.Printf("%s ", curr.Value)
		curr = curr.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println("=== DBMS Ready ===")
	fmt.Println("Примеры команд:")
	fmt.Println("  MPUSH myarr apple")
	fmt.Println("  FPUSH mylist HEAD banana")
	fmt.Println("  SPUSH mystack cherry")
	fmt.Println("  PRINT myarr")
	fmt.Println("Введите EXIT для выхода.")
	fmt.Println("==================")

	reg := NewRegistry()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			break // EOF (Ctrl+D или Ctrl+Z)
		}

		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		if strings.ToUpper(input) == "EXIT" {
			break
		}

		// Разбиваем строку на слова (токены)
		tokens := strings.Fields(input)
		if len(tokens) == 0 {
			continue
		}

		cmd := tokens[0]

		// 1. Обработка команды PRINT
		if cmd == "PRINT" {
			if len(tokens) < 2 {
				fmt.Println("PRINT: missing structure name")
				continue
			}
			name := tokens[1]
			if s, exists := reg.structures[name]; exists {
				s.Print()
			} else {
				fmt.Printf("Structure '%s' not found\n", name)
			}
			continue
		}

		// 2. Обработка команд модификации данных
		if len(tokens) < 2 {
			fmt.Println("Missing structure name")
			continue
		}

		name := tokens[1]
		typeName := cmd[0] // Первая буква команды определяет тип структуры (M, F, L, S, Q)

		s := reg.GetOrCreate(name, typeName)
		if s == nil {
			fmt.Printf("Unknown structure type: %c\n", typeName)
			continue
		}

		// Передаем команде только аргументы (без cmd и name)
		args := tokens[2:]
		s.Execute(cmd, args)
	}
	
	fmt.Println("Exiting...")
}