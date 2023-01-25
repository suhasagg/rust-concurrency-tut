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
