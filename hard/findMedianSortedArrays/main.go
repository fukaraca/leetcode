package main

import "fmt"

func main() {
	nums1 := []int{1, 3, 5, 8, 13, 21}
	nums2 := []int{2, 4, 6, 8, 23, 67}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	//1,2,3,4,5,6,8,8,13,21,23,67 : 7
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	ret := 0.0
	len1 := len(nums1)
	len2 := len(nums2)
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	num1min, num1max, midway := 0, len1, (len1+len2+1)/2
	for {
		maxleft := 0
		minright := 0
		num1half := (num1min + num1max) / 2
		num2half := midway - num1half
		if num1half < len1 && nums2[num2half-1] > nums1[num1half] {
			num1min = num1half + 1
			continue
		} else if num1half > 0 && nums1[num1half-1] > nums2[num2half] {
			num1max = num1half - 1
			continue
		} else {
			if num1half == 0 {
				maxleft = nums2[num2half-1]
			} else if num2half == 0 {
				maxleft = nums1[num1half-1]
			} else {
				maxleft = maks(nums2[num2half-1], nums1[num1half-1])
			}
			if (len1+len2)%2 == 1 {
				return float64(maxleft)
			}

			if num1half == len1 {
				minright = nums2[num2half]
			} else if num2half == len2 {
				minright = nums1[num1half]
			} else {
				minright = min(nums2[num2half], nums1[num1half])
			}
			ret = (float64(maxleft) + float64(minright)) / 2
			break
		}
	}
	return ret
}

func maks(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

func min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	}
	return num1
}
