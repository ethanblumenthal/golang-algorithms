package main

// O(n^2) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	nums := map[int]bool{}
	
	for _, num := range array {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}
