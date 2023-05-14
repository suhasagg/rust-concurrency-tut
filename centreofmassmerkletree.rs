use std::collections::HashMap;

struct MerkleTree {
    nodes: HashMap<usize, MerkleNode>,
}

struct MerkleNode {
    hash: Vec<u8>,
    left: Option<usize>,
    right: Option<usize>,
}

impl MerkleTree {
    fn new() -> MerkleTree {
        MerkleTree {
            nodes: HashMap::new(),
        }
    }

    fn add_leaf(&mut self, hash: Vec<u8>) -> usize {
        let node_id = self.nodes.len();
        self.nodes.insert(node_id, MerkleNode {
            hash,
            left: None,
            right: None,
        });
        node_id
    }

    fn build(&mut self) {
        let mut queue = vec![0];
        while !queue.is_empty() {
            let node_id = queue.pop().unwrap();
            let node = &self.nodes[&node_id];
            if node.left.is_some() && node.right.is_some() {
                let left_hash = self.nodes[&node.left].hash;
                let right_hash = self.nodes[&node.right].hash;
                node.hash = sha256::digest(&left_hash[..] + &right_hash[..]);
            }
            if node.left.is_some() {
                queue.push(node.left.unwrap());
            }
            if node.right.is_some() {
                queue.push(node.right.unwrap());
            }
        }
    }

    fn center_of_mass(&self) -> f64 {
        let mut total_mass = 0.0;
        let mut total_x = 0.0;
        for (_, node) in self.nodes.iter() {
            total_mass += node.hash.len() as f64;
            total_x += node.hash.len() as f64 * node.hash[0] as f64;
        }
        return total_x / total_mass;
    }
}
