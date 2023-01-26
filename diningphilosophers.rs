use std::sync::{Arc, Mutex};
use std::thread;

const N: usize = 5;

fn philosopher(id: usize, left: Arc<Mutex<()>>, right: Arc<Mutex<()>>) {
    for i in 0..3 {
        println!("Philosopher {} is thinking", id);
        let _left = left.lock().unwrap();
        let _right = right.lock().unwrap();
        println!("Philosopher {} is eating", id);
    }
}

fn main() {
    let chopsticks: Vec<Arc<Mutex<()>>> = (0..N).map(|_| Arc::new(Mutex::new(()))).collect();
    let mut handles = vec![];
    for i in 0..N {
        let left = chopsticks[i].clone();
        let right = chopsticks[(i + 1) % N].clone();
        handles.push(thread::spawn(move || {
            philosopher(i, left, right);
        }));
    }
    for handle in handles {
        handle.join().unwrap();
    }
}
