use std::collections::VecDeque;
use std::sync::{Arc, Mutex};
use std::thread;
use std::io::Read;

use std::net::TcpStream;

fn main() {
    let visited = Arc::new(Mutex::new(Vec::new()));
    let seed_urls = vec!["https://www.example.com", "https://www.example2.com"];

    let mut queue = VecDeque::new();
    for url in seed_urls {
        queue.push_back(url);
    }

    let num_threads = 8;
    let mut handles = Vec::new();
    for _ in 0..num_threads {
        let visited = visited.clone();
        let handle = thread::spawn(move || {
            while let Some(url) = queue.pop_front() {
                if !visited.lock().unwrap().contains(&url) {
                    visited.lock().unwrap().push(url);

                    // Parse the url to get the host and request path
                    let (host, path) = match parse_url(url) {
                        Ok((host, path)) => (host, path),
                        Err(err) => {
                            println!("Error: {:?}", err);
                            continue;
                        },
                    };

                    // Connect to the host
                    let mut stream = match TcpStream::connect(&host) {
                        Ok(stream) => stream,
                        Err(err) => {
                            println!("Error: {:?}", err);
                            continue;
                        },
                    };

                    // Build and send the HTTP request
                    let request = format!("GET {} HTTP/1.1\r\nHost: {}\r\n\r\n", path, host);
                    stream.write_all(request.as_bytes()).unwrap();

                    // Read the response
                    let mut response = String::new();
                    stream.read_to_string(&mut response).unwrap();

                    // parse the response and extract new URLs
                    // ...
                    // add new URLs to the queue
                    // ...
                }
            }
        });
        handles.push(handle);
    }

    for handle in handles {
        handle.join().unwrap();
    }
}

// parse_url function
fn parse_url(url: &str) -> Result<(&str, &str), &str> {
    let parts: Vec<&str> = url.split("://").collect();
    if parts.len() != 2 {
        return Err("Invalid URL");
    }

    let host_path: Vec<&str> = parts[1].split("/").collect();
    let host = host_path[0];
    let path = if host_path.len() > 1 {
        host_path[1..].join("/")
    } else {
        ""
    };
    Ok((host, path))
}
