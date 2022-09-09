package main

import (
	"fmt"
)

func main() {

	vals := []int{2, 4, 1}
	_ = test(vals)

}
func test(prices []int) int {
	var buy int = 0
	var sell int = 1
	var maxprofit = 0

	for sell < len(prices) {

		fmt.Println(buy, sell, maxprofit)

		if prices[buy] < prices[sell] {
			profit := prices[sell] - prices[buy]
			maxprofit = findmax(profit, maxprofit)
		} else {
			buy = sell
		}

		sell++
	}
	fmt.Println(maxprofit)
	return maxprofit
}

func findmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
