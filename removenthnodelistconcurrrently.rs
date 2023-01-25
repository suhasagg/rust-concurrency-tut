use std::sync::{Arc, RwLock};
use std::thread;

struct Node {
    next: Option<Arc<RwLock<Node>>>,
    value: i32,
}

struct List {
    head: Option<Arc<RwLock<Node>>>,
}

impl List {
    fn new() -> Self {
        List { head: None }
    }

    fn insert(&mut self, value: i32) {
        let new_node = Arc::new(RwLock::new(Node {
            next: self.head.clone(),
            value,
        }));
        self.head = Some(new_node);
    }

    fn remove_nth(&mut self, n: usize) {
        let mut current = self.head.clone();
        for _ in 0..n {
            if let Some(node) = current {
                current = node.read().unwrap().next.clone();
            }
        }
        if let Some(node) = current {
            let mut node = node.write().unwrap();
            node.next = node.next.as_ref().unwrap().read().unwrap().next.clone();
        }
    }
    fn print_list(&self) {
        let mut current = self.head.clone();
        while let Some(node) = current {
            let node = node.read().unwrap();
            print!("{} ", node.value);
            current = node.next.clone();
        }
        println!();
    }
}

fn main() {
    let mut list = List::new();
    list.insert(1);
    list.insert(2);
    list.insert(3);
    list.insert(4);
    list.insert(5);

    list.print_list(); // 5 4 3 2 1

    let list = Arc::new(RwLock::new(list));
    let list2 = list.clone();

    let handle = thread::spawn(move || {
        let mut list = list2.write().unwrap();
        list.remove_nth(2);
    });

    handle.join().unwrap();
    let list = list.read().unwrap();
    list.print_list(); // 5 4 2
}
