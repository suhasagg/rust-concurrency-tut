func subsets(set []int) [][]int {
    n := len(set)
    result := make([][]int, 0, 1<<uint(n))

    for i := 0; i < (1 << uint(n)); i++ {
        subset := make([]int, 0)
        for j := 0; j < n; j++ {
            if i&(1<<uint(j)) != 0 {
                subset = append(subset, set[j])
            }
        }
        result = append(result, subset)
    }

    return result
}

//In this example, the powerSet function uses an iterative approach to generate the power set of the given set. It uses a for-loop to iterate through all possible subsets represented by a bit mask, and checks if the current subset includes the corresponding element of the set by using the bitwise AND operator. The function then appends the current subset to the result.
//This approach is faster than the recursive approach because it avoids the overhead of creating and destroying multiple stack frames and it's also faster than the concurrent approach because it doesn't need to handle synchronization
//Beats 100 percent soultions in leet code 0 ms runtime
