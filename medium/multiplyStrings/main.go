package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(multiply("999", "999"))
	/*
		4672
		584
	*/
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	l1 := len(num1)
	l2 := len(num2)
	if l2 > l1 {
		num1, num2 = num2, num1
		l1 = len(num1)
		l2 = len(num2)
	}
	sums := [][]uint8{}

	for i := l2 - 1; i >= 0; i-- {
		temp := []uint8{}
		leftover := 0
		for j := l1 - 1; j >= 0; j-- {
			mul := leftover + int(num2[i]-'0')*int(num1[j]-'0')
			temp = append(temp, uint8(mul%10))
			leftover = mul / 10
		}
		if leftover != 0 {
			temp = append(temp, uint8(leftover))
		}
		for k, m := 0, len(temp)-1; k < m; k, m = k+1, m-1 {
			temp[k], temp[m] = temp[m], temp[k]
		}
		for n := 0; n < l2-1-i; n++ {
			temp = append(temp, uint8(0))
		}
		sums = append(sums, temp)
	}
	last := sums[len(sums)-1]

	for m, sum := range sums {
		if m == len(sums)-1 {
			continue
		}
		carry := uint8(0)
		for i, j := len(sum)-1, 0; i >= 0; i, j = i-1, j+1 {
			temp := carry + last[len(last)-1-j] + sum[i]
			carry = temp / 10
			last[len(last)-1-j] = temp % 10
			if i == 0 && carry != 0 {

				if len(last)-1-j-1 < 0 {

					t := []uint8{0}
					t = append(t, last...)
					last = t
				}
				a := len(last) - 1 - j - 1
				for carry > 0 && a >= 0 {
					fmt.Println(last, carry, a)
					last[a] += carry
					carry = last[a] / 10
					a -= 1
				}

			}
		}

	}

	var ret strings.Builder
	for _, u := range last {
		if u > 9 {
			ret.Write([]byte{u - 9 + '0', 0 + '0'})
			continue
		}
		ret.WriteByte(u + '0')
	}
	return ret.String()
}
