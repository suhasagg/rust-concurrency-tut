package main

import (
    "bytes"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

type Trie struct {
    root *Node
}

type Node struct {
    key   []byte
    value []byte
    children map[byte]*Node
    hash []byte
}

func (t *Trie) Insert(key, value []byte) {
    t.root = t.insert(t.root, key, value, 0)
}

func (t *Trie) insert(node *Node, key, value []byte, depth int) *Node {
    if node == nil {
        node = &Node{children: make(map[byte]*Node)}
    }
    if len(key) == depth {
        node.key = make([]byte, len(key))
        copy(node.key, key)
        node.value = make([]byte, len(value))
        copy(node.value, value)
        return node
    }
    c := key[depth]
    node.children[c] = t.insert(node.children[c], key, value, depth+1)
    node.hash = t.hashNode(node)
    return node
}

func (t *Trie) Get(key []byte) []byte {
    node := t.get(t.root, key, 0)
    if node != nil {
        return node.value
    }
    return nil
}

func (t *Trie) get(node *Node, key []byte, depth int) *Node {
    if node == nil {
        return nil
    }
    if len(key) == depth {
        return node
    }
    c := key[depth]
    return t.get(node.children[c], key, depth+1)
}

func (t *Trie) MerkleRoot() string {
    return hex.EncodeToString(t.root.hash)
}

func (t *Trie) hashNode(node *Node) []byte {
    var buffer bytes.Buffer
    if node.value != nil {
        buffer.Write(node.value)
    }
    for _, child := range node.children {
        buffer.Write(child.hash)
    }
    hasher := sha256.New()
    hasher.Write(buffer.Bytes())
    return hasher.Sum(nil)
}

func main() {
    trie := &Trie{}
    trie.Insert([]byte("hello"), []byte("world"))
    trie.Insert([]byte("hell"), []byte("go"))
    fmt.Println(string(trie.Get([]byte("hello"))))
    fmt.Println(string(trie.Get([]byte("hell"))))
    fmt.Println("Merkle Root:", trie.MerkleRoot())
    }
