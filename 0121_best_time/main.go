package main

import (
	"fmt"
	"math"
	"reflect"
)

/*
Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:
	Input: [7,1,5,3,6,4]
	Output: 5
	Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
				 Not 7-1 = 6, as selling price needs to be larger than buying price.

Example 2:
	Input: [7,6,4,3,1]
	Output: 0
	Explanation: In this case, no transaction is done, i.e. max profit = 0.
*/

const UnKnown = math.MaxInt64

func maxProfit(prices []int) (profit int) {
	if len(prices) == 0 {
		return 0
	}

	lowestBuyPrice := UnKnown
	for _, price := range prices {
		if lowestBuyPrice != UnKnown {
			nowProfit := price - lowestBuyPrice
			if profit < nowProfit {
				profit = nowProfit
			}
		}

		if price < lowestBuyPrice {
			lowestBuyPrice = price
		}
	}

	return profit
}

// test case
func main() {
	result := maxProfit([]int{7, 1, 5, 3, 6, 4})
	if reflect.DeepEqual(result, 5) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = maxProfit([]int{7, 6, 4, 3, 1})
	if reflect.DeepEqual(result, 0) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
