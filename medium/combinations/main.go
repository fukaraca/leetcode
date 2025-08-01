package main

func main() {
}

func combine(n int, k int) [][]int {
	out := make([][]int, 0)

	var dfs func(i int, curr []int)
	dfs = func(i int, curr []int) {
		if len(curr) == k {
			out = append(out, append([]int{}, curr...))
			return
		}
		for ; i <= n; i++ {
			curr = append(curr, i)
			dfs(i+1, curr)
			curr = curr[:len(curr)-1]
		}
	}

	dfs(1, make([]int, 0))
	return out
}
