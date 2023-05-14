package main

import (
    "fmt"
    "math"
)

type Node struct {
    Value int
    Left  *Node
    Right *Node
}

func main() {
    // Create a binary tree.
    root := &Node{
        Value: 1,
        Left:  &Node{
            Value: 2,
            Left:  &Node{
                Value: 3,
                Left:  nil,
                Right: nil,
            },
            Right: &Node{
                Value: 4,
                Left:  nil,
                Right: nil,
            },
        },
        Right: &Node{
            Value: 5,
            Left:  &Node{
                Value: 6,
                Left:  nil,
                Right: nil,
            },
            Right: &Node{
                Value: 7,
                Left:  nil,
                Right: nil,
            },
        },
    }

    // Convert the binary tree to vectors.
    vectors := convertNodeToVector(root)

    // Calculate the cosine distance between two nodes.
    cos := calculateCosineDistance(vectors[0], vectors[1])

    // Print the cosine distance.
    fmt.Println(cos)
}

func convertNodeToVector(node *Node) []float64 {
    if node == nil {
        return nil
    }
    return []float64{node.Value} + convertNodeToVector(node.Left) + convertNodeToVector(node.Right)
}

func calculateCosineDistance(vector1 []float64, vector2 []float64) float64 {
    dotProduct := 0.0
    for i := range vector1 {
        dotProduct += vector1[i] * vector2[i]
    }
    magnitude1 := 0.0
    for i := range vector1 {
        magnitude1 += vector1[i] * vector1[i]
    }
    magnitude1 = math.Sqrt(magnitude1)
    magnitude2 := 0.0
    for i := range vector2 {
        magnitude2 += vector2[i] * vector2[i]
    }
    magnitude2 = math.Sqrt(magnitude2)
    return dotProduct / (magnitude1 * magnitude2)
}
