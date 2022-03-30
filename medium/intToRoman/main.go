package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(intToRoman(3999))
}

func intToRoman(num int) string {
	str := strconv.Itoa(num)
	size := len(str)
	ret := ""
	for i, d := range str {
		switch {
		case size-i == 4:
			for j := int32(0); j < (d - 48); j++ {
				ret += "M"
			}
		case size-i == 3:
			if d >= 48 && d < 52 {
				for j := int32(0); j < (d - 48); j++ {
					ret += "C"
				}
			} else if d == 52 {
				ret += "CD"
			} else if d == 53 {
				ret += "D"
			} else if d > 53 && d < 57 {
				ret += "D"
				for j := int32(0); j < (d - 53); j++ {
					ret += "C"
				}
			} else {
				ret += "CM"
			}
		case size-i == 2:
			if d >= 48 && d < 52 {
				for j := int32(0); j < (d - 48); j++ {
					ret += "X"
				}
			} else if d == 52 {
				ret += "XL"
			} else if d == 53 {
				ret += "L"
			} else if d > 53 && d < 57 {
				ret += "L"
				for j := int32(0); j < (d - 53); j++ {
					ret += "X"
				}
			} else {
				ret += "XC"
			}
		case size-i == 1:
			if d >= 48 && d < 52 {
				for j := int32(0); j < (d - 48); j++ {
					ret += "I"
				}
			} else if d == 52 {
				ret += "IV"
			} else if d == 53 {
				ret += "V"
			} else if d > 53 && d < 57 {
				ret += "V"
				for j := int32(0); j < (d - 53); j++ {
					ret += "I"
				}
			} else {
				ret += "IX"
			}

		}
	}
	return ret
}
