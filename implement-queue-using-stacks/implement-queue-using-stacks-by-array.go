package main

import (
	"errors"
	"fmt"
)

type SliceStack struct {
	arr []int
	stackSize int
}

func (p *SliceStack) isEmpty() bool {
	return p.stackSize == 0
}

func (p *SliceStack) Size() int {
	return p.stackSize
}

func (p *SliceStack) Top() int {
	if (p.isEmpty()) {
		panic(errors.New("stack is empty"))
	}

	return p.arr[p.stackSize - 1]
}

func (p *SliceStack) Pop() int {
	if p.stackSize <= 0 {
		panic(errors.New("stack is empty"))
	}

	p.stackSize--
	ret := p.arr[p.stackSize]
	p.arr = p.arr[:p.stackSize]
	return ret
}

func (p *SliceStack) Push(element int) {
	p.arr = append(p.arr, element)
	p.stackSize = p.stackSize + 1
}

func main() {
	SliceMode()
}

func SliceMode() {
	defer func ()  {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	sliceStack := &SliceStack{arr: make([]int, 0)}

	sliceStack.Push(1)
	fmt.Println("栈顶元素为: ", sliceStack.Top())
	fmt.Println("栈大小为: ", sliceStack.Size())
	sliceStack.Pop()
	fmt.Println("after pop: ", sliceStack.Size())
	sliceStack.Pop()
}