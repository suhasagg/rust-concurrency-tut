import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type VerkleNode struct {
	value    byte
	children []*VerkleNode
}

func newVerkleNode(value byte) *VerkleNode {
	return &VerkleNode{
		value:    value,
		children: make([]*VerkleNode, 256),
	}
}

func makeVerkleTree(strs []string) *VerkleNode {
	root := newVerkleNode(0)
	for _, s := range strs {
		binaryStr := toBinaryString(s)
		node := root
		for i := 0; i < len(binaryStr); i++ {
			b := binaryStr[i]
			if node.children[b] == nil {
				node.children[b] = newVerkleNode(b)
			}
			node = node.children[b]
		}
		node.children[0] = newVerkleNode(0) // add terminator node
	}
	return root
}

func toBinaryString(s string) string {
	var buf bytes.Buffer
	for _, c := range []byte(s) {
		binary.Write(&buf, binary.LittleEndian, c)
	}
	return buf.String()
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	root := makeVerkleTree(strs)
	var prefix bytes.Buffer
	node := root
	for {
		var child *VerkleNode
		for i := 0; i < 256; i++ {
			if node.children[i] != nil {
				if child == nil {
					child = node.children[i]
				} else {
					return prefix.String()
				}
			}
		}
		if child == nil {
			return prefix.String()
		}
		if child.value == 0 {
			return prefix.String()
		}
		prefix.WriteByte(child.value)
		node = child
	}
}
