package main

import (
	"fmt"
	"math"
)

/*

Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

Example:
	Input: [2,3,1,1,4]
	Output: 2
	Explanation: The minimum number of jumps to reach the last index is 2.
    			 Jump 1 step from index 0 to 1, then 3 steps to the last index.

Note:
	You can assume that you can always reach the last index.
*/

const UnReachable = math.MaxInt64

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	topPosition := len(nums) - 1
	numberOfJumps := make([]int, len(nums), len(nums))
	for i := topPosition - 1; i >= 0; i-- {
		numberOfJumps[i] = UnReachable
		maxReachPosition := i + nums[i]

		// can jump to the top position directly
		if maxReachPosition >= topPosition {
			numberOfJumps[i] = 1
			continue
		}

		// reach top position by pivot
		for pivot := i + 1; pivot <= maxReachPosition; pivot++ {
			if numberOfJumps[pivot] == UnReachable {
				continue
			}
			if numberOfJumps[pivot] == 1 {
				numberOfJumps[i] = numberOfJumps[pivot] + 1
				break
			}

			if (numberOfJumps[pivot] + 1) < numberOfJumps[i] {
				numberOfJumps[i] = numberOfJumps[pivot] + 1
			}
		}
	}

	return numberOfJumps[0]
}

// test case
func main() {
	result := jump([]int{2, 3, 1, 1, 4})
	if result != 2 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

	result = jump([]int{1})
	if result != 0 {
		panic(fmt.Sprintf("test case fail, result: %+v", result))
	}

}
