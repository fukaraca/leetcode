package main

import "fmt"

func main() {
	fmt.Println(canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	ls := make([][]int, numCourses)
	for i := range ls {
		ls[i] = make([]int, 0)
	}
	for _, pr := range prerequisites {
		ls[pr[0]] = append(ls[pr[0]], pr[1])
	}
	v := make(map[int]uint8) // 0 not visited 1 visiting 2 visited
	var dfs func(dep int) bool
	dfs = func(dep int) bool {
		switch v[dep] {
		case 2:
			return true
		case 1:
			return false // found cycle
		}

		check := ls[dep]
		if len(check) == 0 {
			v[dep] = 2
			return true
		}
		v[dep] = 1
		for _, c := range check {
			if !dfs(c) {
				return false
			}
		}
		v[dep] = 2
		return true
	}

	for i, l := range ls {
		if len(l) == 0 {
			continue
		}
		if !dfs(i) {
			return false
		}
		//v = make(map[int]uint8)
	}
	return true
}
