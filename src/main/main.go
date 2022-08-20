package main

import (
	"fmt"
	"strconv"
	"unsafe"
 )


func constructMaximumBinaryTree(nums []int) *TreeNode {
	tree := make([]*TreeNode, len(nums))
	var stk []int
	for i, num := range nums {
		tree[i] = &TreeNode{Val: num}
		for len(stk) > 0 && num > nums[stk[len(stk)-1]] {
			tree[i].Left = tree[stk[len(stk)-1]]
			stk = stk[:len(stk)-1]
		}
		if len(stk) > 0 {
			tree[stk[len(stk)-1]].Right = tree[i]
		}
		stk = append(stk, i)
	}
	return tree[stk[0]]
}

func main() {
	fmt.Println(ParseTreeNode("[1,2,3,null,4,null]"))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func StringToBytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&str))
}

func ListToString(c []interface{}) (res string) {
	res = "["
	for _, o := range c {
		switch v := o.(type) {
		case string:
			res += "\"" + v + "\""
		case *int:
			if v == nil {
				res += "null"
			} else {
				res += strconv.Itoa(*v)
			}
		}
		res += ","
	}
	if len(res) > 1 {
		res = res[:len(res)-1]
	}
	return res + "]"
}

