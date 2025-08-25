package main

func main() {
}

func longestSubarray(nums []int) int {
	var pre int
	var max, count int
	var del, once bool
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			once = true
			if !del {
				if count > 0 {
					del = true
					pre = i - 1
					continue
				}
				pre = i
				continue
			} else {
				max = maxi(max, count)
				count = 0
				i = pre
				del = false
			}
		}

	}
	var res = maxi(max, count)
	if !once && res != 0 {
		res--
	}

	return res
}

func maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}
