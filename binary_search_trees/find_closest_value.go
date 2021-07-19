package main

type BST struct {
	Value int
	Left  *BST
	Right *BST
}


// Average: O(log(n)) time | O(log(n)) space
func (tree *BST) FindClosestValue(target int) int {
	return tree.findClosestValue(target, tree.Value)
}

func (tree *BST) findClosestValue(target int, closest int) int {
	if absoluteDifference(target, closest) > absoluteDifference(target, tree.Value) {
		closest = tree.Value
	}
	if target < tree.Value && tree.Left != nil {
		return tree.Left.findClosestValue(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findClosestValue(target, closest)
	}
	return closest
}

func absoluteDifference(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}