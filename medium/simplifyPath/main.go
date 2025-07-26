package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "/.../a/../b/c////../d/./"
	fmt.Println(simplifyPath(s))
	fmt.Println(simplifyPath("/"))
}

func simplifyPath(path string) string {
	tokens := strings.Split(path, "/")
	// .. , sondaki/,
	var validToken []string
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case ".", "":
		case "..":
			if len(validToken) > 0 {
				validToken = validToken[:len(validToken)-1]
			}
		default:
			validToken = append(validToken, tokens[i])
		}
	}

	return "/" + strings.Join(validToken, "/")
}
