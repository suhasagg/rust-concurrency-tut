package main

import (
	"container/heap"
	"fmt"
)

type Message struct {
	Priority int
	Content  string
}

type PriorityQueue []*Message

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority > pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Message)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

func main() {
	highPriority := make(chan Message)
	mediumPriority := make(chan Message)
	lowPriority := make(chan Message)

	go func() {
		highPriority <- Message{Priority: 3, Content: "High Priority Message"}
		mediumPriority <- Message{Priority: 2, Content: "Medium Priority Message"}
		lowPriority <- Message{Priority: 1, Content: "Low Priority Message"}
	}()

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	for i := 0; i < 3; i++ {
		select {
		case msg := <-highPriority:
			heap.Push(&pq, &msg)
		case msg := <-mediumPriority:
			heap.Push(&pq, &msg)
		case msg := <-lowPriority:
			heap.Push(&pq, &msg)
		}
	}

	for pq.Len() > 0 {
		msg := heap.Pop(&pq).(*Message)
		fmt.Printf("Priority: %d, Content: %s\n", msg.Priority, msg.Content)
	}
}
