use std::sync::{Arc, Mutex, Condvar};
use std::thread;

struct BoundedBlockingQueue {
    size: usize,
    queue: Vec<i32>,
    lock: Mutex<()>,
    not_empty: Condvar,
    not_full: Condvar,
}

impl BoundedBlockingQueue {
    fn new(size: usize) -> Arc<Self> {
        let queue = Arc::new(BoundedBlockingQueue {
            size,
            queue: Vec::new(),
            lock: Mutex::new(()),
            not_empty: Condvar::new(),
            not_full: Condvar::new(),
        });

        queue
    }

    fn put(&self, element: i32) {
        let mut guard = self.lock.lock().unwrap();
        while self.queue.len() == self.size {
            guard = self.not_full.wait(guard).unwrap();
        }
        self.queue.push(element);
        self.not_empty.notify_one();
    }

    fn take(&self) -> i32 {
        let mut guard = self.lock.lock().unwrap();
        while self.queue.is_empty() {
            guard = self.not_empty.wait(guard).unwrap();
        }
        let element = self.queue.remove(0);
        self.not_full.notify_one();
        element
    }
}

fn main() {
    let queue = BoundedBlockingQueue::new(2);
    queue.put(1);
    queue.put(2);
    println!("{}", queue.take());
    println!("{}", queue.take());
}
