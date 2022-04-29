package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	a := "foobarthebarmanfoothebar"
	words := []string{"foo", "bar", "the"}
	fmt.Println(findSubstring(a, words))
	/*
		"wordgoodgoodgoodbestword"
		["word","good","best","good"]
	*/

}

func findSubstring(s string, words []string) []int {
	lenB := len(words[0])
	lenW := len(words) * lenB
	var sumW int32
	for _, word := range words {
		for _, i3 := range word {
			sumW += i3
		}
	}
	ret := []int{}
	if len(s) < lenW {
		return ret
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i] > words[j]
	})
	for i := 0; i < len(s)-lenW+1; i++ {
		tempSum := int32(0)
		tempSum = 0
		for j := i; j < len(s) && j < i+lenW; j++ {
			tempSum += int32(s[j])
		}
		if tempSum == sumW {
			//check here and return append indices if OK
			temp := []string{}
			for k := i; k < i+lenW; k += lenB {
				temp = append(temp, s[k:k+lenB])
			}
			sort.Slice(temp, func(i1, j1 int) bool {
				return temp[i1] > temp[j1]
			})
			if reflect.DeepEqual(temp, words) {
				ret = append(ret, i)
			}
		}
	}
	return ret
}
