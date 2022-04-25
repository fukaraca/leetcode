package main

import "fmt"

func main() {
	a := "LALL"
	fmt.Println(checkRecord(a))
}

func checkRecord(s string) bool {
	absent := 0
	late := 0
	for _, r := range s {
		switch {
		case r == 'A':
			absent++
			if absent >= 2 {
				return false
			}
		case r == 'L':
			late++
			if late >= 3 {
				return false
			}
		default:
			late = 0
		}
	}
	return true

}
