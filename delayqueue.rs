use std::collections::BinaryHeap;
use std::time::{Instant, Duration};

struct Item {
    timestamp: Instant,
    value: String,
}

impl Item {
    fn new(value: String, delay: Duration) -> Self {
        Self {
            timestamp: Instant::now() + delay,
            value,
        }
    }
}

impl PartialEq for Item {
    fn eq(&self, other: &Self) -> bool {
        self.timestamp == other.timestamp
    }
}

impl Eq for Item {}

impl PartialOrd for Item {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        self.timestamp.partial_cmp(&other.timestamp)
    }
}

impl Ord for Item {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.timestamp.cmp(&other.timestamp)
    }
}

fn main() {
    let delay = Duration::from_secs(2);
    let mut q = BinaryHeap::new();

    std::thread::spawn(move || loop {
        while let Some(item) = q.pop() {
            if item.timestamp <= Instant::now() {
                println!("{}", item.value);
            } else {
                q.push(item);
                break;
            }
        }
        std::thread::sleep(delay);
    });

    q.push(Item::new("Hello, world!".to_string(), Duration::from_secs(3)));
}
