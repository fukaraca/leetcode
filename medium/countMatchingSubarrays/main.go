package main

func main() {
}

func countMatchingSubarrays(nums []int, pattern []int) int {
	var count int

	var fn func(ind int) bool

	fn = func(ind int) bool {
		for i := 0; i < len(pattern); i++ {
			switch pattern[i] {
			case 1:
				if !(nums[ind] > nums[ind-1]) {
					return false
				}
			case 0:
				if nums[ind] != nums[ind-1] {
					return false
				}
			case -1:
				if !(nums[ind] < nums[ind-1]) {
					return false
				}
			}
			ind++
		}
		return true
	}

	for i := 1; i < len(nums)-len(pattern)+1; i++ {
		if fn(i) {
			count++
		}
	}
	return count
}
