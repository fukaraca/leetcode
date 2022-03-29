package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(myAtoi("mil -500 mil"))
}
func myAtoi(s string) int {

	ret := ""
	check := true
	for _, l := range s {
		if !(l == 45 || l == 32 || l == 43 || (l > 47 && l < 58)) && check {
			return 0
		} else if (l == 45 || l == 43 || (l > 47 && l < 58)) && check {
			check = false
			ret += string(l)
		} else if !check && (l > 47 && l < 58) {
			ret += string(l)
		} else if !check && !(l > 47 && l < 58) {
			break
		}
	}
	intret, _ := strconv.ParseInt(ret, 10, 32)
	/*	if err != nil {
		panic(err)
	}*/

	return int(intret)
}
