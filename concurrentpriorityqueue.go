type PriorityQueue struct {
    data   []int
    ch     chan int
    mutex  sync.Mutex
    update chan bool
}

func (pq *PriorityQueue) Insert(val int) {
    pq.mutex.Lock()
    defer pq.mutex.Unlock()
    pq.data = append(pq.data, val)
    pq.update <- true
    go pq.heapifyUp(len(pq.data) - 1)
}

func (pq *PriorityQueue) Remove() int {
    pq.mutex.Lock()
    defer pq.mutex.Unlock()
    if len(pq.data) == 0 {
        return -1
    }
    val := pq.data[0]
    pq.data[0] = pq.data[len(pq.data)-1]
    pq.data = pq.data[:len(pq.data)-1]
    pq.update <- true
    go pq.heapifyDown(0)
    return val
}

func (pq *PriorityQueue) heapifyUp(i int) {
    for i > 0 && pq.data[i] < pq.data[(i-1)/2] {
        pq.data[i], pq.data[(i-1)/2] = pq.data[(i-1)/2], pq.data[i]
        i = (i - 1) / 2
    }
}

func (pq *PriorityQueue) heapifyDown(i int) {
    for {
        minIdx := i
        left := 2*i + 1
        right := 2*i + 2
        if left < len(pq.data) && pq.data[left] < pq.data[minIdx] {
            minIdx = left
        }
        if right < len(pq.data) && pq.data[right] < pq.data[minIdx] {
            minIdx = right
        }
        if minIdx == i {
            break
        }
        pq.data[i], pq.data[minIdx] = pq.data[minIdx], pq.data[i]
        i = minIdx
    }
}
