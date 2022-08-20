package main

import (
	"encoding/json"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (t TreeNode) String() string {
	li := make([]*int, 0)
	q := []*TreeNode{&t}
	for len(q) > 0 {
		now := q[0]
		q = q[1:]
		if now == nil {
			li = append(li, nil)
		} else {
			li = append(li, &now.Val)
			q = append(q, now.Left)
			q = append(q, now.Right)
		}
	}
	for i := len(li) - 1; i > 0; i-- {
		if li[i] == nil {
			li = li[:i]
		} else {
			break
		}
	}
	fuck := "["
	for _, num := range li {
		if num != nil {
			fuck += strconv.Itoa(*num)
		} else {
			fuck += "null"
		}
		fuck += ","
	}
	if len(fuck) > 1 {
		fuck = fuck[:len(fuck)-1]
	}
	return fuck + "]"
}

func ParseTreeNode(s string) (res *TreeNode) {
	var arr []*int
	err := json.Unmarshal(StringToBytes(s), &arr)
	if err != nil {
		return
	}
	n := len(arr)
	if n == 0 || arr[0] == nil {
		return
	}
	res = &TreeNode{Val: *arr[0]}
	q := []*TreeNode{0: res}
	for i := 1; i < n; i++ {
		now := q[0]
		q = q[1:]
		if arr[i] != nil {
			now.Left = &TreeNode{Val: *arr[i]}
			q = append(q, now.Left)
		}
		i++
		if i == n {
			break
		}
		if arr[i] != nil {
			now.Right = &TreeNode{Val: *arr[i]}
			q = append(q, now.Right)
		}
	}
	return
}
