package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flower", "flower", "flower"}))
}

func longestCommonPrefix(strs []string) string {
	ret := ""
	if strs[0] == "" {
		return ""
	}
	check := true
	for i := 0; check; i++ {
		pre := strs[0][:i+1]
		for _, str := range strs {
			if str == "" {
				return ""
			}
			if str[:i+1] != pre {
				return ret
			}
			if len(str)-1 == i {
				check = false
			}
		}
		ret = pre
	}
	return ret
}
