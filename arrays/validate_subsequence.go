package main

// O(n) time | O(1) space
func IsValidSubsequence(array []int, sequence []int) bool {
	seqIdx := 0
	
	for _, val := range array {
		if seqIdx == len(sequence) {
			break
		}
		
		if val == sequence[seqIdx] {
			seqIdx += 1
		}
	}
	return seqIdx == len(sequence)
}
