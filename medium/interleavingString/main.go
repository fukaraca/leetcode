package main

func main() {
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != (len(s1) + len(s2)) {
		return false
	}
	type a struct {
		i1, i2 int
	}
	m := make(map[a]bool)
	l1, l2, l3 := len(s1), len(s2), len(s3)
	var fn func(i1, i2, i3 int) bool

	fn = func(i1, i2, i3 int) bool {
		if l3 == i3 {
			return true
		}
		if v, ok := m[a{i1, i2}]; ok {
			return v
		}
		var found bool
		switch {
		case i1 < l1 && s1[i1] == s3[i3] && i2 < l2 && s2[i2] == s3[i3]:
			found = fn(i1+1, i2, i3+1) || fn(i1, i2+1, i3+1)
		case i1 < l1 && s1[i1] == s3[i3]:
			found = fn(i1+1, i2, i3+1)
		case i2 < l2 && s2[i2] == s3[i3]:
			found = fn(i1, i2+1, i3+1)
		}

		m[a{i1, i2}] = found
		return found
	}

	return fn(0, 0, 0)
}
