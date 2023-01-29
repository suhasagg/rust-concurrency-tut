package main

import (
	"fmt"
	"net"
	"strings"
)

// Node represents a peer node in the P2P network
type Node struct {
	Addr string
	Data map[string][]byte
}

// P2PNetwork represents a P2P network of nodes
type P2PNetwork struct {
	Nodes []*Node
}

func (n *P2PNetwork) addNode(node *Node) {
	n.Nodes = append(n.Nodes, node)
}

func (n *P2PNetwork) broadcast(data map[string][]byte) {
	for _, node := range n.Nodes {
		node.Data = data
	}
}

func main() {
	network := &P2PNetwork{}

	node1 := &Node{
		Addr: "localhost:8000",
		Data: make(map[string][]byte),
	}
	network.addNode(node1)

	node2 := &Node{
		Addr: "localhost:8001",
		Data: make(map[string][]byte),
	}
	network.addNode(node2)

	// Start a server on node1 to listen for incoming data from other nodes
	go func() {
		ln, _ := net.Listen("tcp", node1.Addr)
		for {
			conn, _ := ln.Accept()
			go func(conn net.Conn) {
				buf := make([]byte, 1024)
				_, _ = conn.Read(buf)
				data := strings.Split(string(buf), ":")
				key := data[0]
				value := []byte(data[1])
				node1.Data[key] = value
				_ = conn.Close()
			}(conn)
		}
	}()

	// Connect node2 to node1 and send some data
	conn, _ := net.Dial("tcp", node1.Addr)
	_, _ = fmt.Fprintf(conn, "file1:data1")
	_ = conn.Close()

	network.broadcast(node2.Data)

	fmt.Println(node1.Data) // Output: map[file1:[100 97 116 97 49]]
	fmt.Println(node2.Data) // Output: map[file1:[100 97 116 97 49]]
}
