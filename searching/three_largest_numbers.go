package main

import "math"

// O(n) time | O(1) space
func FindThreeLargestNumbers(array []int) []int {
	threeLargest := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	
	for _, val := range array {
		updateLargest(threeLargest, val)
	}
	return threeLargest
}

func updateLargest(threeLargest []int, val int) {
	if val > threeLargest[2]  {
		shiftAndUpdate(threeLargest, val, 2)
	} else if val > threeLargest[1] {
		shiftAndUpdate(threeLargest, val, 1)
	} else if val > threeLargest[0] {
		shiftAndUpdate(threeLargest, val, 0)
	}
}

func shiftAndUpdate(array []int, val int, idx int) {
	for i := 0; i < idx + 1; i++ {
		if i == idx {
			array[i] = val
		} else {
			array[i] = array[i+1]
		}
	}
}