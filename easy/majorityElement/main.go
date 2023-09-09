package main

import "fmt"

func main() {
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // return 2

}
func majorityElement(nums []int) int {
	st := struct {
		max, am int
		buckets map[int]int
	}{buckets: make(map[int]int, len(nums))}
	for _, num := range nums {
		st.buckets[num]++
		// t:=st.buckets[num]
		if st.buckets[num] > st.am {
			st.max = num
			st.am = st.buckets[num]
		}
	}
	return st.max
}
