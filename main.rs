use std::sync::{ Arc, Mutex, mpsc::channel };
use std::thread;

use std::time::Duration;

fn main() {
    
    // test_move();
    test_channel();
    // test_receive1();
    test_mutex();
    test_mutex1();
}

pub fn test_mutex1() {
    let counter = Arc::new(Mutex::new(0));
    let mut handles = vec![];

    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        
        let handle = thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }

    println!("Result : {}", *counter.lock().unwrap());
}

pub fn test_mutex() {
    let x = Mutex::new(6);
    
    {
       
        let mut num = x.lock().unwrap();
        *num = 10;
    }
    println!("x = {:?}", x);
}

pub fn test_receive1() {
    let (tx, rx) = channel();
    let tx1 = tx.clone();

    thread::spawn(move || {
        let strings = vec![
            String::from("1:hi"),
            String::from("1:hello"),
            String::from("1:good"),
            String::from("1:china"),
        ];

        for item in strings {
            tx.send(item).unwrap();
            thread::sleep(Duration::from_millis(200));
        }
    });
    
    thread::spawn(move || {
        let strings = vec![
            String::from("hi"),
            String::from("hello"),
            String::from("good"),
            String::from("china"),
        ];

        for item in strings {
            tx1.send(item).unwrap();
            thread::sleep(Duration::from_millis(200));
        }
    });

    for result in rx {
        println!("{:?}", result);
    }
}

pub fn test_receive() {
    let (tx, rx) = channel();
    thread::spawn(move || {
        let strings = vec![
            String::from("hi"),
            String::from("hello"),
            String::from("good"),
            String::from("china"),
        ];

        for item in strings {
            tx.send(item).unwrap();
            thread::sleep(Duration::from_millis(200));
        }
    });

   
    for result in rx {
        println!("{:?}", result);
    }
}

pub fn test_channel() {
    let (sender, receiver) = channel();

    thread::spawn(move || {
        let message = String::from("Hello");
        sender.send(message).unwrap();
        // borrow of moved value: `message` value borrowed here after move
        // println!("{:?}", message);
    });

   
    let result = receiver.recv().unwrap();
    println!("{:?}", result);
}

pub fn test_move() {
    let vec = vec![1, 2, 3, 4, 5];
   
    let hadnle = thread::spawn(move || { 
        println!("{:?}", vec);
    });
    hadnle.join().unwrap();
}

pub fn test_spawn() {
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1));
        }
    });

    
    handle.join().unwrap();

    for i in 1..5 {
        println!("hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1));
    }
}
