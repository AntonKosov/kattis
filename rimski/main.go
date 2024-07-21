package main

import (
	"fmt"
	"sort"
)

func main() {
	var num string
	fmt.Scanf("%s", &num)
	fmt.Println(findMin(num))
}

func findMin(num string) string {
	return nums[sortDigits(num)]
}

var nums map[string]string

func init() {
	nums = make(map[string]string, len(tens)*len(ones))
	for _, t := range tens {
		for _, o := range ones {
			num := t + o
			sortedNum := sortDigits(num)
			if _, ok := nums[sortedNum]; !ok {
				nums[sortedNum] = num
			}
		}
	}
}

var (
	tens = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ones = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func sortDigits(num string) string {
	b := []byte(num)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
