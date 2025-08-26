package main

func main() {
}

func canCompleteCircuit(gas []int, cost []int) int {
	var cand int
	var totalCost int
	var tank int

	for i := 0; i < len(gas); i++ {
		totalCost += cost[i] - gas[i]
		tank += gas[i]
		if tank < cost[i] {
			cand = i + 1
			tank = 0
			continue
		}
		tank -= cost[i]
	}
	if totalCost > 0 {
		return -1
	}
	return cand
}
