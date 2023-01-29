package main

import (
	"crypto/sha1"
	"fmt"
	"net"
)

const (
	IDLength = 20
	BucketSize = 8
)

type Node struct {
	ID   [IDLength]byte
	Addr net.UDPAddr
}

type RoutingTable struct {
	Node     Node
	Buckets [IDLength * 8]map[string]Node
}

func (rt *RoutingTable) AddNode(n Node) {
	bucketIndex := rt.getBucketIndex(n.ID)
	if rt.Buckets[bucketIndex] == nil {
		rt.Buckets[bucketIndex] = make(map[string]Node)
	}
	rt.Buckets[bucketIndex][string(n.ID[:])] = n
}

func (rt *RoutingTable) FindClosest(targetID [IDLength]byte, count int) []Node {
	var closest []Node
	for i := 0; i < IDLength*8; i++ {
		if rt.Buckets[i] == nil {
			continue
		}
		for _, node := range rt.Buckets[i] {
			closest = append(closest, node)
			if len(closest) >= count {
				return closest
			}
		}
	}
	return closest
}

func (rt *RoutingTable) getBucketIndex(id [IDLength]byte) int {
	for i := 0; i < IDLength; i++ {
		for j := 0; j < 8; j++ {
			if (id[i]>>uint8(7-j))&0x01 != 0 {
				return i*8 + j
			}
		}
	}
	return IDLength*8 - 1
}

func main() {
	var rt RoutingTable

	// Generate a fake node ID and address for testing
	nodeID := sha1.Sum([]byte("node1"))
	nodeAddr, _ := net.ResolveUDPAddr("udp", "localhost:8080")

	rt.Node = Node{ID: nodeID, Addr: *nodeAddr}

	// Add some fake nodes to the routing table
	rt.AddNode(Node{ID: sha1.Sum([]byte("node2")), Addr: net.UDPAddr{}})
	rt.AddNode(Node{ID: sha1.Sum([]byte("node3")), Addr: net.UDPAddr{}})
	rt.AddNode(Node{ID: sha1.Sum([]byte("node4")), Addr: net.UDPAddr{}})

	// Find closest nodes to a fake target ID
	closest := rt.FindClosest(sha1.Sum([]byte("target")), 2)
	}
