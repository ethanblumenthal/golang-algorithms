package main

import (
	"sort"
)

// O(nlogn) time | O(n) space - where n is the length of the input array
func SortedSquaredArray(array []int) []int {
	squaredArray := make([]int, len(array))
	
	for idx, val := range array {
		squaredArray[idx] = val * val
	}
	
	sort.Ints(squaredArray)
	return squaredArray
}
