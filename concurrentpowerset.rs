use std::sync::mpsc;
use std::thread;

fn power_set(set: &[i32], idx: usize, current: &[i32], tx: mpsc::Sender<Vec<i32>>) {
    if idx == set.len() {
        tx.send(current.to_vec()).unwrap();
    } else {
        let (tx1, rx1) = mpsc::channel();
        thread::spawn(move || power_set(set, idx + 1, &[&current[..], &[set[idx]]].concat(), tx1));

        let (tx2, rx2) = mpsc::channel();
        thread::spawn(move || power_set(set, idx + 1, current, tx2));
        for res in rx1.iter().chain(rx2.iter()) {
            tx.send(res).unwrap();
        }
    }
}

fn main() {
    let set = vec![1, 2, 3];
    let (tx, rx) = mpsc::channel();
    thread::spawn(move || power_set(&set, 0, &[], tx));
    for subset in rx.iter() {
        println!("{:?}", subset);
    }
}


//In this example, the power_set function is called concurrently using the thread::spawn method. The power_set function generates the power set of the given set using backtracking, where the idx parameter represents the current position in the set and the current parameter represents the current subset. The tx parameter is used to send the subsets to the main thread.
//The power_set function makes a choice to either include or exclude the current element and recursively calls itself for the next element.
//The mpsc::channel is used to create channels that allow communication between different threads.
//It's important to note that this approach might not be the most efficient for large sets, and other approaches should be considered
