package main

func main() {
}

func merge(nums1 []int, m int, nums2 []int, n int) {

	p1 := m - 1
	p2 := n - 1
	p3 := m + n - 1
	for ; p1 >= 0 && p2 >= 0; p3-- {
		if nums2[p2] >= nums1[p1] {
			nums1[p3] = nums2[p2]
			p2--
			continue
		}
		nums1[p3] = nums1[p1]
		p1--
	}
	if p2 >= 0 {
		copy(nums1[:p3+1], nums2[:p2+1])
	}

}
