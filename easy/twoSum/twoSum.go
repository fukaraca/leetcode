package twoSum

func twoSum(nums []int, target int) []int {
	newMap:=make(map[int]int)
	for i,v:=range nums{
		complement:=target-v
		if j,ok:=newMap[complement];ok{
			return []int{j,i}
		}
		newMap[v]=i
	}
	return nil
}
