package main

import (
    "fmt"
    "math"
    "sort"
)

type Node struct {
    Data int
    Left *Node
    Right *Node
}

func main() {
    // Create a binary tree.
    root := &Node{Data: 1}
    root.Left = &Node{Data: 2}
    root.Right = &Node{Data: 3}
    root.Left.Left = &Node{Data: 4}
    root.Left.Right = &Node{Data: 5}
    root.Right.Left = &Node{Data: 6}
    root.Right.Right = &Node{Data: 7}

    // Calculate the angles between the nodes.
    angles := make([]float64, 0, len(root.Children()))
    for _, child := range root.Children() {
        angles = append(angles, math.Atan2(child.Y-root.Y, child.X-root.X))
    }

    // Sort the angles.
    sort.Float64s(angles)

    // Split the tree at the levels where the angles are equal.
    for i := 1; i < len(angles); i++ {
        if angles[i] == angles[i-1] {
            // Split the tree at this level.
            node := root.Children()[i]
            root.Left = node
            root.Right = nil
            root = node
        }
    }

    // Print the optimized tree.
    fmt.Println(root)
}
