package main

import "fmt"

var groups map[int]int
var cache map[int]int

func init() {
	groups = map[int]int{}
	cache = map[int]int{}
}

func main() {
	max, cows, days := read()
	for _, c := range cows {
		groups[c]++
	}
	calc(maxDay(days), max)
	for _, day := range days {
		fmt.Println(cache[day])
	}
}

func calc(maxDay, max int) {
	cache[0] = sum()
	for d := 1; d <= maxDay; d++ {
		nextGroup := map[int]int{}
		totalFarms := 0
		for cows, farms := range groups {
			cows *= 2
			if cows > max {
				sep := cows / 2
				cows -= sep
				nextGroup[sep] += farms
				totalFarms += farms
			}
			nextGroup[cows] += farms
			totalFarms += farms
		}

		groups = nextGroup
		cache[d] = totalFarms
	}
}

func sum() int {
	s := 0
	for _, farms := range groups {
		s += farms
	}
	return s
}

func maxDay(days []int) int {
	m := 0
	for _, day := range days {
		if day > m {
			m = day
		}
	}
	return m
}

func read() (max int, cows, days []int) {
	var n, m int
	fmt.Scanf("%d %d %d", &max, &n, &m)
	for i := 0; i < n; i++ {
		var f int
		fmt.Scanf("%d", &f)
		cows = append(cows, f)
	}
	for i := 0; i < m; i++ {
		var d int
		fmt.Scanf("%d", &d)
		days = append(days, d)
	}
	return max, cows, days
}
