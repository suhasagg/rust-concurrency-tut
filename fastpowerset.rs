fn power_set(set: &[i32]) -> Vec<Vec<i32>> {
    let n = set.len();
    let mut result = Vec::with_capacity(1 << n);

    for i in 0..(1 << n) {
        let mut subset = Vec::new();
        for j in 0..n {
            if (i & (1 << j)) != 0 {
                subset.push(set[j]);
            }
        }
        result.push(subset);
    }

    result
}

fn main() {
    let set = vec![1, 2, 3];
    let result = power_set(&set);
    println!("{:?}", result);
    // Output: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]
}
