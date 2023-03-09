import (
    "crypto/sha256"
    "encoding/hex"
)

type Trie struct {
    root *node
}

type node struct {
    children map[byte]*node
    value    []byte
    hash     []byte
}

// Initialize the trie
func NewTrie() *Trie {
    return &Trie{root: &node{children: make(map[byte]*node)}}
}

// Insert a key-value pair into the trie
func (t *Trie) Insert(key, value []byte) {
    t.root = t.insertHelper(t.root, key, value, 0)
}

func (t *Trie) insertHelper(current *node, key, value []byte, depth int) *node {
    if current == nil {
        current = &node{children: make(map[byte]*node)}
    }

    if depth == len(key) {
        current.value = value
        current.hash = t.calculateHash(current)
        return current
    }

    current.children[key[depth]] = t.insertHelper(current.children[key[depth]], key, value, depth+1)
    current.hash = t.calculateHash(current)
    return current
}

// Search for a value in the trie based on a key
func (t *Trie) Get(key []byte) []byte {
    current := t.getHelper(t.root, key, 0)
    if current != nil {
        return current.value
    }
    return nil
}

func (t *Trie) getHelper(current *node, key []byte, depth int) *node {
    if current == nil {
        return nil
    }

    if depth == len(key) {
        return current
    }

    return t.getHelper(current.children[key[depth]], key, depth+1)
}

// Calculate the hash of a node based on its children and value
func (t *Trie) calculateHash(n *node) []byte {
    if len(n.children) == 0 && n.value == nil {
        return nil
    }

    hashData := make([][]byte, 0)
    for k, v := range n.children {
        hashData = append(hashData, []byte{k}, v.hash)
    }
    if n.value != nil {
        hashData = append(hashData, n.value)
    }

    concatenated := bytes.Join(hashData, []byte{})
    hash := sha256.Sum256(concatenated)
    return hash[:]
}

// Check if a key exists in the trie
func (t *Trie) Contains(key []byte) bool {
    return t.Get(key) != nil
}

// Check if a prefix exists in the trie
func (t *Trie) StartsWith(prefix []byte) bool {
    current := t.root
    for _, b := range prefix {
        if _, ok := current.children[b]; !ok {
            return false
        }
        current = current.children[b]
    }
    return true
}
