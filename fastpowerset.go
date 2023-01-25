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
