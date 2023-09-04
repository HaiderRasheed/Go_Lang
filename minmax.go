package main

import (
	"math"
)

type MinStack struct {
	stack []int
	min   int
}

func NewMinStack() MinStack {
	return MinStack{min: math.MaxInt64}
}
func (ms *MinStack) Push(item int) {
	ms.stack = append(ms.stack, item)
	if item < ms.min {
		ms.min = item
	}
}
func (ms *MinStack) Pop() {
	length := len(ms.stack)
	last := ms.stack[length-1]
	ms.stack = ms.stack[:length-1]
	if last != ms.min {
		return
	}
	ms.min = math.MaxInt64
	for _, x := range ms.stack {
		if x < ms.min {
			ms.min = x
		}
	}
}
func (ms *MinStack) GetTop() int {
	return ms.stack[len(ms.stack)-1]
}
func (ms *MinStack) GetMin() int {
	return ms.min
}
