const wordSize = 32

type Trie struct {
    children [wordSize]*Trie
    value    interface{}
    isEnd    bool
}

func (t *Trie) Insert(key uint32, value interface{}) {
    curr := t
    for i := 0; i < wordSize; i++ {
        if curr.children[key&1] == nil {
            curr.children[key&1] = &Trie{}
        }
        curr = curr.children[key&1]
        key >>= 1
    }
    curr.value = value
    curr.isEnd = true
}

func (t *Trie) Search(key uint32) (interface{}, bool) {
    curr := t
    for i := 0; i < wordSize; i++ {
        if curr.children[key&1] == nil {
            return nil, false
        }
        curr = curr.children[key&1]
        key >>= 1
    }
    if !curr.isEnd {
        return nil, false
    }
    return curr.value, true
}
