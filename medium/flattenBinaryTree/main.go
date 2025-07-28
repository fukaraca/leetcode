package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "1,2,5,3,4,null,6"
	root := GetNode(s)
	flatten(root)
	fmt.Println(root)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	var st []*TreeNode
	t := root
	var dfs func(r *TreeNode)
	dfs = func(r *TreeNode) {
		if r == nil {
			return
		}
		st = append(st, r)
		dfs(r.Left)
		dfs(r.Right)
	}
	dfs(root)
	for i := 1; i < len(st); i++ {
		t.Left = nil
		t.Right = st[i]
		t = t.Right
	}
}

func GetNode(s string) *TreeNode {
	sVals := strings.Split(s, ",")
	vals := make([]*int, len(sVals))
	for i := range sVals {
		if sVals[i] != "null" {
			t, _ := strconv.Atoi(sVals[i])
			vals[i] = &t
		}
	}
	l := len(vals)
	if l == 0 || vals[0] == nil {
		return nil
	}
	var i int
	ch := make(chan *TreeNode, len(vals))
	out := &TreeNode{Val: *vals[i]}
	i++
	ch <- out
	for i < l {
		curr := <-ch
		if vals[i] != nil {
			ct := &TreeNode{Val: *vals[i]}
			curr.Left = ct
			ch <- ct
		}
		i++
		if i < l && vals[i] != nil {
			ct := &TreeNode{Val: *vals[i]}
			curr.Right = ct
			ch <- ct
		}
		i++
	}
	return out
}
