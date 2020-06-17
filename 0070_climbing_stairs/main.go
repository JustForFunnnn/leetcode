package main

import (
	"fmt"
	"reflect"
)

/*
You are climbing a stair case. It takes n steps to reach to the top.

Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?

Note: Given n will be a positive integer.
*/

func climbStairs(n int) int {
	resultCache := make([]int, n, n)
	return waysToTopStairs(0, n, resultCache)
}

func waysToTopStairs(now, top int, resultCache []int) (ways int) {
	if now == top {
		return 1
	}

	if top < now {
		return 0
	}

	if resultCache[now] != 0 {
		return resultCache[now]
	}

	// 1 step
	ways += waysToTopStairs(now+1, top, resultCache)
	// 2 steps
	ways += waysToTopStairs(now+2, top, resultCache)

	// cache result
	resultCache[now] = ways
	return ways
}

// test case
func main() {
	result := climbStairs(2)
	if reflect.DeepEqual(result, 2) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = climbStairs(3)
	if reflect.DeepEqual(result, 3) != true {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}
}
