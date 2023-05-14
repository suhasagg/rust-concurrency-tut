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

    fn gravitational_force_of_attraction(&self, node1_id: usize, node2_id: usize) -> f64 {
        let node1 = &self.nodes[&node1_id];
        let node2 = &self.nodes[&node2_id];
        let mass1 = node1.hash.len() as f64;
        let mass2 = node2.hash.len() as f64;
        let distance = self.distance(node1_id, node2_id);
        return G * mass1 * mass2 / distance.powi(2);
    }

    fn distance(&self, node1_id: usize, node2_id: usize) -> f64 {
        let mut distance = 0.0;
        let mut node1 = node1_id;
        let mut node2 = node2_id;
        while node1 != node2 {
            let level1 = self.level(node1);
            let level2 = self.level(node2);
            if level1 > level2 {
                let diff = level1 - level2;
                for _ in 0..diff {
                    node1 = self.nodes[&node1].left.unwrap();
                }
            } else {
                let diff = level2 - level1;
                for _ in 0..diff {
                    node2 = self.nodes[&node2].left.unwrap();
                }
            }
            distance += 1.0;
        }
        return distance;
    }

    fn level(&self, node_id: usize) -> usize {
        let mut level = 0;
        let mut node = node_id;
        while node != 0 {
            node = self.nodes[&node].parent.unwrap();
            level += 1;
        }
        return level;
    }
}
