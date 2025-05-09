package main

import (
	"fmt"
)

func main() {
	recipes := []string{"ju", "fzjnm", "x", "e", "zpmcz", "h", "q"}
	ingredients := [][]string{
		{"d"},
		{"hveml", "f", "cpivl"},
		{"cpivl", "zpmcz", "h", "e", "fzjnm", "ju"},
		{"cpivl", "hveml", "zpmcz", "ju", "h"},
		{"h", "fzjnm", "e", "q", "x"},
		{"d", "hveml", "cpivl", "q", "zpmcz", "ju", "e", "x"},
		{"f", "hveml", "cpivl"}}
	supplies := []string{"f", "hveml", "cpivl", "d"}
	fmt.Println(findAllRecipes(recipes, ingredients, supplies))
}

func findAllRecipes(recipes []string, ingredients [][]string, supplies []string) []string {
	var out []string
	supp := make(map[string]bool)
	recps := make(map[string]int)
	visited := make(map[int]bool)
	for i := range supplies {
		supp[supplies[i]] = true
	}
	for i := range recipes {
		recps[recipes[i]] = i
	}

	var checkRecipe func(int) bool
	checkRecipe = func(i int) bool {
		if visited[i] {
			return false
		}
		visited[i] = true
		for j := range ingredients[i] {
			if !supp[ingredients[i][j]] {
				if v, ok := recps[ingredients[i][j]]; !ok || v == -1 {
					return false
				} else if !checkRecipe(v) {
					recps[ingredients[i][j]] = -1
					return false
				}
				supp[ingredients[i][j]] = true
			}
		}
		return true
	}
	for i := range recipes {
		if checkRecipe(i) {
			out = append(out, recipes[i])
		}
		visited = make(map[int]bool)
	}
	return out
}
