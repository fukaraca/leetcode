package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(countAndSay(4))
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	ret := strings.Builder{}
	//ret:=""
	say := countAndSay(n - 1)
	counts := make(map[int32]int)
	temp := int32(say[0])
	for i, s := range say {
		if s == temp {
			counts[s-48]++
		}
		if i != len(say)-1 && say[i+1] != say[i] {
			ret.WriteString(strconv.Itoa(counts[s-48]))
			ret.WriteString(strconv.Itoa(int(s) - 48))
			//ret = fmt.Sprintf("%s%d%d", ret, counts[s-48], s-48)
			temp = int32(say[i+1])
			counts = make(map[int32]int)
		} else if i == len(say)-1 {
			ret.WriteString(strconv.Itoa(counts[s-48]))
			ret.WriteString(strconv.Itoa(int(s) - 48))

		}
	}
	return ret.String()
}
