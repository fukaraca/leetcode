package lengthOfLongestSubstring

func lengthOfLongestSubstring(s string) int {
	start := 0
	max := 1
	substr := make(map[int32]int)
	if s == "" {
		return 0
	}
	for i, l := range s {
		if ind, ok := substr[l]; ok {
			if i-start > max {
				max = i - start
			}
			if ind >= start {
				start = ind + 1
			}
			substr[l] = i
		} else {
			substr[l] = i
		}

	}
	if len(s)-start > max {
		max = len(s) - start
	}
	return max
}
