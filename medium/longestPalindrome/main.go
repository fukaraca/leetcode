package longestPalindrome

func longestPalindrome(s string) string {
	leng := len(s)
	longest := 0
	res := 0
	if leng == 1 {
		return s
	}
	for i := 0; i < leng; i++ {
		counter := 0
		for j, k := i, leng-1; k >= j; k-- {
			if s[k] == s[j] && (k != j) {
				j++
				counter++
				counter++
			} else {
				if k == j {
					counter++
					j++
				} else {
					if counter > 0 {
						k += counter / 2
						counter = 0
						j = i
					} else {
						counter = 0
						j = i
					}
				}
			}
		}
		if counter > longest {
			res = i
			longest = counter
		}
	}
	return s[res : res+longest]
}
