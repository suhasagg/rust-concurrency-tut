import (
    "fmt"
    "sync"
)

// powerSet generates the power set of the given set using backtracking
func powerSet(set []int, index int, current []int, result *[][]int, wg *sync.WaitGroup) {
    if index == len(set) {
        // add the current subset to the result
        subset := make([]int, len(current))
        copy(subset, current)
        *result = append(*result, subset)
    } else {
        // make a choice to include the current element
        current = append(current, set[index])
        wg.Add(1)
        go powerSet(set, index+1, current, result, wg)
        current = current[:len(current)-1]

        // make a choice to exclude the current element
        wg.Add(1)
        go powerSet(set, index+1, current, result, wg)
    }
    wg.Done()
}

func main() {
    set := []int{1, 2, 3}
    var result [][]int
    var wg sync.WaitGroup
    wg.Add(1)
    go powerSet(set, 0, []int{}, &result, &wg)
    wg.Wait()
    fmt.Println(result)
    // Output: [[] [1] [2] [1 2] [3] [1 3] [2 3] [1 2 3]]
}
