package main

import (
    "fmt"
    "sync"
)

// BoundedBlockingQueue is a struct that implements a bounded blocking queue
type BoundedBlockingQueue struct {
    size int
    queue []interface{}
    mutex sync.Mutex
    notEmpty sync.Cond
    notFull sync.Cond
}

// NewBoundedBlockingQueue returns a new instance of BoundedBlockingQueue
func NewBoundedBlockingQueue(size int) *BoundedBlockingQueue {
    queue := &BoundedBlockingQueue{size: size}
    queue.notEmpty = sync.Cond{L: &queue.mutex}
    queue.notFull = sync.Cond{L: &queue.mutex}
    return queue
}

// Put adds an element to the queue
func (queue *BoundedBlockingQueue) Put(element interface{}) {
    queue.mutex.Lock()
    defer queue.mutex.Unlock()
    for queue.size == len(queue.queue) {
        queue.notFull.Wait()
    }
    queue.queue = append(queue.queue, element)
    queue.notEmpty.Signal()
}

// Take removes and returns the first element from the queue
func (queue *BoundedBlockingQueue) Take() interface{} {
    queue.mutex.Lock()
    defer queue.mutex.Unlock()
    for len(queue.queue) == 0 {
        queue.notEmpty.Wait()
    }
    element := queue.queue[0]
    queue.queue = queue.queue[1:]
    queue.notFull.Signal()
    return element
}

func main() {
    queue := NewBoundedBlockingQueue(2)
    queue.Put(1)
    queue.Put(2)
    fmt.Println(queue.Take())
    fmt.Println(queue.Take())
}
