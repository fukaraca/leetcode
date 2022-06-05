package runningSum

func runningSum(nums []int) []int {
	ret := []int{}
	for i, _ := range nums {
		temp := 0
		for j := 0; j < i+1; j++ {
			temp += nums[j]
		}
		ret = append(ret, temp)
	}
	return ret
}
