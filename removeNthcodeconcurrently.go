package main

import (
    "fmt"
    "sync"
)

// Node represents a node in a linked list
type Node struct {
    next *Node
    value int
}

// List represents a linked list
type List struct {
    head *Node
    mu sync.RWMutex
}

// Insert adds a new node to the list
func (l *List) Insert(value int) {
    l.mu.Lock()
    defer l.mu.Unlock()

    newNode := &Node{value: value}
    newNode.next = l.head
    l.head = newNode
}

// RemoveNth removes the nth node from the list
func (l *List) RemoveNth(n int) {
    var current *Node
    l.mu.RLock()
    if n == 0 {
        l.head = l.head.next
        l.mu.RUnlock()
        return
    }
    current = l.head
    for i := 1; i < n; i++ {
        current = current.next
    }
    l.mu.RUnlock()
    l.mu.Lock()
    current.next = current.next.next
    l.mu.Unlock()
}

// PrintList prints the list
func (l *List) PrintList() {
    l.mu.RLock()
    defer l.mu.RUnlock()

    current := l.head
    for current != nil {
        fmt.Print(current.value, " ")
        current = current.next
    }
    fmt.Println()
}

func main() {
    list := &List{}
    list.Insert(1)
    list.Insert(2)
    list.Insert(3)
    list.Insert(4)
    list.Insert(5)

    list.PrintList() // prints 5 4 3 2 1

    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        list.RemoveNth(2)
    }()
    wg.Wait()
    list.PrintList() // prints 5 4 2 1
}
