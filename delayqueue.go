package main

import (
    "container/heap"
    "time"
)

type Item struct {
    Timestamp time.Time
    Value     interface{}
}

type ItemHeap []*Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].Timestamp.Before(h[j].Timestamp) }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *ItemHeap) Push(x interface{}) {
    *h = append(*h, x.(*Item))
}

func (h *ItemHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    delay := 2 * time.Second
    q := &ItemHeap{}
    heap.Init(q)

    go func() {
        for {
            now := time.Now()
            for q.Len() > 0 {
                item := heap.Pop(q).(*Item)
                if item.Timestamp.Before(now) {
                    // Process item.Value
                } else {
                    heap.Push(q, item)
                    break
                }
            }
            time.Sleep(delay)
        }
    }()

    heap.Push(q, &Item{
        Timestamp: time.Now().Add(3 * time.Second),
        Value:     "Hello, world!",
    })
}
