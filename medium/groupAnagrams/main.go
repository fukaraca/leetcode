package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(groupAnagrams(nil))
}

func groupAnagrams(strs []string) [][]string {
	grMap := make(map[string][]string)
	for _, str := range strs {
		temp := []byte(str)
		sort.Slice(temp, func(i, j int) bool {
			return temp[i] < temp[j]
		})
		grMap[string(temp)] = append(grMap[string(temp)], str)
	}
	ret := [][]string{}
	for _, st := range grMap {
		ret = append(ret, st)
	}
	return ret

}
