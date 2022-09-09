package main

import (
	"fmt"
)

func main() {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("%v", result[0])
}
func twoSum(nums []int, target int) []int {
	//https://leetcode.com/problems/two-sum/

	m := make(map[int]int)
	for index, num := range nums {
		if val, ok := m[num]; ok {
			return []int{val, index}
		} else {
			m[target-num] = index
		}
	}
	return []int{-1, -1}
}
