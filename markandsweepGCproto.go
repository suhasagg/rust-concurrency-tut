package main

import (
    "fmt"
    "sync"
)

type Object struct {
    value int
    next  *Object
}

var objects []*Object
var mu sync.Mutex

func markAndSweep() {
    // Stop the world
    mu.Lock()
    defer mu.Unlock()

    // Mark all reachable objects
    markReachableObjects()

    // Sweep through all objects and reclaim the garbage
    for _, obj := range objects {
        if !obj.reachable {
            obj = nil
        }
    }

    // Perform memory compaction
    compactMemory()
}

func markReachableObjects() {
    // Mark the global objects as reachable
    for _, obj := range objects {
        obj.reachable = true
    }

    // Mark all objects reachable from the global objects
    for _, obj := range objects {
        markObject(obj)
    }
}

func markObject(obj *Object) {
    if obj == nil || obj.reachable {
        return
    }

    obj.reachable = true

    markObject(obj.next)
}

func compactMemory() {
    // Move all reachable objects to the front of the slice
    for i := 0; i < len(objects); i++ {
        if !objects[i].reachable {
            for j := i + 1; j < len(objects); j++ {
                if objects[j].reachable {
                    objects[i], objects[j] = objects[j], objects[i]
                    break
                }
            }
        }
    }

    // Resize the slice to remove the garbage
    objects = objects[:len(objects)-countGarbage()]
}

func countGarbage() int {
    count := 0
    for _, obj := range objects {
        if !obj.reachable {
            count++
        }
    }
    return count
}

func main() {
    // Create some objects
    objects = append(objects, &Object{value: 1})
    objects = append(objects, &Object{value: 2, next: objects[0]})
    objects = append(objects, &Object{value: 3, next: objects[1]})

    // Print the objects
    for _, obj := range objects {
        fmt.Println(obj.value)
    }

    // Run the mark-and-sweep algorithm
    markAndSweep()

    // Print the objects again
    for _, obj := range objects {
        fmt.Println(obj.value)
    }
}
