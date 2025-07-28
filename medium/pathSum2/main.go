package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "-260,-202,-903,-980,-570,-858,218,764,-300,205,null,-35,null,null,-204,950,-769,258,-652,614,-584,76,817,-192,null,null,-114,880,null,-200,71,671,344,801,193,-18,876,-920,-730,222,679,null,-680,null,null,null,-859,744,-261,692,null,-341,-163,null,null,482,-979,205,null,146,165,801,100,-656,714,-629,995,474,307,-581,-150,-941,null,null,null,-937,-69,-23,82,null,-139,-591,null,-453,-861,-370,null,null,null,216,233,null,430,null,5,-110,null,null,-660,624,-510,-588,null,null,381,null,368,559,null,521,-301,null,522,379,null,null,null,null,456,519,null,null,482,349,null,null,19,null,null,288,-811,null,-372,null,null,-536,null,-404,-457,-740,860,null,null,-636,null,null,342,-874,-462,-504,781,855,-392,null,null,null,406,null,-758,541,null,-947,null,null,null,null,null,-964,null,600,-45,null,null,null,null,null,null,null,null,null,-194,null,null,null,-802,null,null,null,-3,null,-792,672,643,null,14,null,null,489,457,null,null,null,null,412,null,558,null,null,null,null,-846,158,-146,null,null,-76,-650,null,-782,null,-127,null,-678,null,null,null,null,null,null,-464,-426,null,-366,null,null,null,null,null,81,-607,716,null,null,-213,null,379,null,null,null,null,644,445,null,null,-419,-845,-720,null,null,-915,null,null,null,null,null,null,-686,594,-243,null,496,null,907,null,null,null,null,null,null,579,873,702,null,null,null,-834,null,null,null,null,null,-300,-214,-466,null,null,972,null,null,null,814,null,-940,null,763,null,null,null,null,-449,-844,null,null,null,null,-47"
	s2 := "-50,100,-500,0,-100,null,null"
	root := GetNode(s)
	//	fmt.Println(root)
	fmt.Println(pathSum(root, -243))
	fmt.Println(pathSum(GetNode(s2), -50))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var out [][]int
	var dfs func(r *TreeNode, sum int, route []int)
	dfs = func(r *TreeNode, sum int, route []int) {
		if r == nil {
			return
		}
		sum += r.Val
		route = append(route, r.Val)
		if r.Left == nil && r.Right == nil && sum == targetSum {
			downstream := make([]int, len(route))
			copy(downstream, route)
			out = append(out, downstream)
			return
		}
		dfs(r.Left, sum, route)
		dfs(r.Right, sum, route)
	}
	dfs(root, 0, make([]int, 0))
	return out
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
