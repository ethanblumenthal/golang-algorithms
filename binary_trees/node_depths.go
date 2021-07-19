package main

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

// O(n) time | O(h) space - where n is the number of nodes in
// the binary tree and h is the height of the binary tree
func NodeDepths(root *BinaryTree) int {
	return calculateNodeDepths(root, 0)
}

func calculateNodeDepths(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + calculateNodeDepths(node.Left, depth+1) + calculateNodeDepths(node.Right, depth+1)
}