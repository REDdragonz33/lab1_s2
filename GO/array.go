package main

import (
	"fmt"
	"strings"
)

type Array struct {
	data []string
}

func NewArray() *Array { return &Array{} }

func (a *Array) GetType() byte { return 'M' }

func (a *Array) Execute(cmd string, args []string) {
	switch cmd {
	case "MPUSH":
		if len(args) == 1 {
			a.data = append(a.data, args[0])
		} else if len(args) == 2 {
			idx, ok := ParseInt(args[0])
			if ok && idx >= 0 && idx <= len(a.data) {
				a.data = append(a.data, "")
				copy(a.data[idx+1:], a.data[idx:])
				a.data[idx] = args[1]
			} else {
				fmt.Println("Index out of range")
			}
		}
	case "MDEL":
		if len(args) < 1 { return }
		idx, ok := ParseInt(args[0])
		if ok && idx >= 0 && idx < len(a.data) {
			a.data = append(a.data[:idx], a.data[idx+1:]...)
		} else {
			fmt.Println("Index out of range")
		}
	case "MGET":
		if len(args) < 1 { return }
		idx, ok := ParseInt(args[0])
		if ok && idx >= 0 && idx < len(a.data) {
			fmt.Println(a.data[idx])
		} else {
			fmt.Println("ERROR")
		}
	case "MSET":
		if len(args) < 2 { return }
		idx, ok := ParseInt(args[0])
		if ok && idx >= 0 && idx < len(a.data) {
			a.data[idx] = args[1]
		} else {
			fmt.Println("Index out of range")
		}
	case "MLEN":
		fmt.Println(len(a.data))
	}
}

func (a *Array) Print() {
	fmt.Printf("Array[%d]: ", len(a.data))
	for _, v := range a.data {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func (a *Array) Save() string {
	return fmt.Sprintf("%d %s", len(a.data), strings.Join(a.data, " "))
}

func (a *Array) Load(data string) {
	a.data = nil
	parts := strings.Fields(data)
	if len(parts) == 0 { return }
	var count int
	fmt.Sscanf(parts[0], "%d", &count)
	for i := 1; i <= count && i < len(parts); i++ {
		a.data = append(a.data, parts[i])
	}
}