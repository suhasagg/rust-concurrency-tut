package main

import (
    "fmt"
    "sync"
)

const N = 5

func philosopher(id int, left, right *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        fmt.Printf("Philosopher %d is thinking\n", id)
        left.Lock()
        right.Lock()
        fmt.Printf("Philosopher %d is eating\n", id)
        right.Unlock()
        left.Unlock()
    }
}

func main() {
    chopsticks := make([]sync.Mutex, N)
    var wg sync.WaitGroup
    wg.Add(N)
    for i := 0; i < N; i++ {
        go philosopher(i, &chopsticks[i], &chopsticks[(i+1)%N], &wg)
    }
    wg.Wait()
}
