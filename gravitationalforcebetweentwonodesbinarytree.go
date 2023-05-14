package main

import (
	"fmt"
)

// Node struct represents a single node in the binary tree
type Node struct {
	Weight int
	Left   *Node
	Right  *Node
}

// BinaryTree represents the entire tree
type BinaryTree struct {
	root *Node
}

// insert method for BinaryTree, to add a new node to the tree
func (t *BinaryTree) insert(data int) *BinaryTree {
	if t.root == nil {
		t.root = &Node{Weight: data, Left: nil, Right: nil}
	} else {
		t.root.insert(data)
	}
	return t
}

// insert method for a Node, to add a new node as a child of the current node
func (n *Node) insert(data int) {
	if n == nil {
		return
	} else if data <= n.Weight {
		if n.Left == nil {
			n.Left = &Node{Weight: data, Left: nil, Right: nil}
		} else {
			n.Left.insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Weight: data, Left: nil, Right: nil}
		} else {
			n.Right.insert(data)
		}
	}
}

// findLCA method for a Node, to find the Lowest Common Ancestor of two nodes
func (n *Node) findLCA(n1, n2 int) *Node {
	if n == nil {
		return nil
	}
	if n.Weight > n1 && n.Weight > n2 {
		return n.Left.findLCA(n1, n2)
	}
	if n.Weight < n1 && n.Weight < n2 {
		return n.Right.findLCA(n1, n2)
	}
	return n
}

// findLevel is a function to find the level of a node in the tree
func findLevel(n *Node, k, level int) int {
	if n == nil {
		return -1
	}
	if n.Weight == k {
		return level
	}

	l := findLevel(n.Left, k, level+1)
	if l != -1 {
		return l
	}
	return findLevel(n.Right, k, level+1)
}

// findDistance is a function to find the distance between two nodes
func findDistance(root *Node, n1, n2 int) int {
	lca := root.findLCA(n1, n2)

	l1 := findLevel(lca, n1, 0)
	l2 := findLevel(lca, n2, 0)

	return l1 + l2
}

// GravitationalForce calculates the gravitational force between two nodes
func GravitationalForce(root *Node, node1, node2 int) float64 {
	distance := float64(findDistance(root, node1, node2))
	return 1 * float64(node1*node2) / (distance * distance)
}

// main function to demonstrate usage of the binary tree
func main() {
	// create a new binary tree and add some nodes
	t := &BinaryTree{}
	t.insert(20).insert(8).insert(22).insert(4).insert(12).insert(10).insert(14)

	// calculate and print the "gravitational force" between two nodes
	fmt.Println(GravitationalForce(t.root, 4, 14))
}
