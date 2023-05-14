package main

import "fmt"

// Node represents a node in the binary tree
type Node struct {
	Value float64  // Value represents the weight of the node
	Left  *Node    // Left represents the left child of the node
	Right *Node    // Right represents the right child of the node
}

// CenterOfMass calculates the total "mass" (defined as the sum of the weighted depths of the nodes)
// and total weight of the subtree rooted at 'node'.
// This function is a helper function that uses depth-first traversal to visit all nodes in the subtree.
func CenterOfMass(node *Node, depth float64) (float64, float64) {
	if node == nil {
		return 0, 0
	}

	// Recursively calculate the total mass and weight of the left and right subtrees
	leftMass, leftWeight := CenterOfMass(node.Left, depth+1)
	rightMass, rightWeight := CenterOfMass(node.Right, depth+1)

	// Calculate the total weight of this subtree: the weight of the node plus the total weights of its subtrees
	totalWeight := leftWeight + rightWeight + node.Value

	// Calculate the total mass of this subtree: the mass of the node (weight times depth) plus the total masses of its subtrees
	totalMass := leftMass + rightMass + depth*node.Value

	return totalMass, totalWeight
}

// ComputeCenterOfMass calculates the "center of mass" of a binary tree rooted at 'root'.
// The "center of mass" is defined as the weighted average of the node depths, where the weights are the node values.
func ComputeCenterOfMass(root *Node) float64 {
	totalMass, totalWeight := CenterOfMass(root, 0)
	return totalMass / totalWeight
}

func main() {
	// Creating a binary tree for demonstration purposes
	root := &Node{
		Value: 1,
		Left: &Node{
			Value: 2,
			Left: &Node{
				Value: 4,
			},
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 3,
		},
	}

	fmt.Println("Center of Mass: ", ComputeCenterOfMass(root))
}
