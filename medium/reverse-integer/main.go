package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(reverse(-50050500))
}

func reverse(x int) int {
	pol := -1
	if x == 0 {
		return 0
	} else if x < 0 {
		x = -x
	} else {
		pol = 1
	}

	str := strconv.Itoa(x)

	sb := strings.Builder{}
	for i := len(str) - 1; i >= 0; i-- {
		sb.WriteString(string(str[i]))
	}
	sayi, err := strconv.Atoi(sb.String())
	if err != nil {
		panic(err)
	}
	if (pol < 0 && sayi > (2147483648)) || (pol > 0 && sayi > 2147483647) {
		return 0
	}

	return pol * sayi
}
