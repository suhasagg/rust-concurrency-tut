use std::sync::{ Arc, RwLock };
use std::thread;

fn main() {
    //Here we have an immutable array with mutable interior cells
    //The length of the array can't change, but internal values can
    //We use the Send sync variants for threading
    let data = Arc::new(vec![
        RwLock::new(1), 
        RwLock::new(2), 
        RwLock::new(3)
    ]);
    
    println!("{:?}", data);
    
    //We spawn a thread for each value, acquire a lock and mutate it
    let handles: Vec<thread::JoinHandle<()>> = (0..3)
        .map(|i| {
            //We need to get ownership of the data here
            //This is so it doesn't try to move the data from the outer scope
            //In the thread, we acquire a write lock guard to the cell and update
            //The lock is released when the guard goes out of scope
            let data = data.clone();
            thread::spawn(move || {
                let mut datum = data[i].write().unwrap();
                *datum += 1;
            })
        })
        .collect();
        
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("{:?}", data);
}

