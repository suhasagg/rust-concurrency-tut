package main

import (
    "sync"
)

type TrieNode struct {
    children map[rune]*TrieNode
    mutex    sync.RWMutex
    isEnd    bool
}

type Trie struct {
    root *TrieNode
}

func (t *Trie) Insert(word string) {
    t.root.mutex.Lock()
    defer t.root.mutex.Unlock()

    curr := t.root
    for _, char := range word {
        if curr.children[char] == nil {
            curr.children[char] = &TrieNode{children: make(map[rune]*TrieNode)}
        }
        curr = curr.children[char]
    }
    curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
    t.root.mutex.RLock()
    defer t.root.mutex.RUnlock()

    curr := t.root
    for _, char := range word {
        curr = curr.children[char]
        if curr == nil {
            return false
        }
    }
    return curr.isEnd
}

func main() {
    trie := Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}

    trie.Insert("hello")
    trie.Insert("suhas")

    fmt.Println(trie.Search("hello")) // Output: true
    fmt.Println(trie.Search("suhas")) // Output: true
    fmt.Println(trie.Search("suha"))  // Output: false
}
