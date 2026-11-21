package main

import "errors"

type Array struct {
	data []string
}

func NewArray() *Array {
	return &Array{data: make([]string, 0)}
}

func (a *Array) PushEnd(val string) {
	a.data = append(a.data, val)
}

func (a *Array) PushAt(idx int, val string) error {
	if idx < 0 || idx > len(a.data) {
		return errors.New("index out of range")
	}
	a.data = append(a.data, "") // Увеличиваем длину
	copy(a.data[idx+1:], a.data[idx:]) // Сдвигаем вправо
	a.data[idx] = val
	return nil
}

func (a *Array) RemoveAt(idx int) error {
	if idx < 0 || idx >= len(a.data) {
		return errors.New("index out of range")
	}
	a.data = append(a.data[:idx], a.data[idx+1:]...) // Сдвигаем влево
	return nil
}

func (a *Array) Get(idx int) (string, error) {
	if idx < 0 || idx >= len(a.data) {
		return "", errors.New("index out of range")
	}
	return a.data[idx], nil
}

func (a *Array) Set(idx int, val string) error {
	if idx < 0 || idx >= len(a.data) {
		return errors.New("index out of range")
	}
	a.data[idx] = val
	return nil
}

func (a *Array) Length() int {
	return len(a.data)
}

func (a *Array) Print() {
	print("Array: ")
	for _, v := range a.data {
		print(v, " ")
	}
	println()
}