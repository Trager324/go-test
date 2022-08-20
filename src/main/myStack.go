package main

import "errors"

type MyStack struct {
	arr []int
	top int
}

var ERR_NO_SUCH_ELEMTNT error = errors.New("no such element")

func New() *MyStack {
	return &MyStack{
		arr: make([]int, 0),
		top: 0,
	}
}

func (it *MyStack)Push(x int) {
	it.arr = append(it.arr, x)
	it.top++
}

func (it *MyStack)Peek() (res int, err error) {
	if it.top == 0 {
		return 0, ERR_NO_SUCH_ELEMTNT
	}
	return it.arr[it.top - 1], nil
}

func (it *MyStack)Pop() (res int, err error) {
	res, err = it.Peek()
	if err != nil {
		return
	}
	it.top--
	it.arr = it.arr[0:it.top]
	return
}

