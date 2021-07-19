package main

type SpecialArray []interface{}

func ProductSum(array []interface{}) int {
	return helper(array, 1)
}

// O(n) time | O(d) space - where n is the total number of elements in the array
// including sub-elements, and d is the greatest depth of special arrays in the array
func helper(array SpecialArray, multiplier int) int {
	sum := 0
	
	for _, val := range array {
		if cast, ok := val.(SpecialArray); ok {
			sum += helper(cast, multiplier + 1)
		} else if cast, ok := val.(int); ok {
			sum += cast
		}
	}
	return sum * multiplier
}