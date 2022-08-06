package main

import (
	"hash/crc32"
	"sort"
	"sync"
)

// Ring is a network of distributed nodes
type Ring struct {
	Nodes Nodes
	rwm   sync.RWMutex
}

// NewRing ...
func NewRing() *Ring {
	return &Ring{
		Nodes: Nodes{},
	}
}

// NewNode ...
func NewNode(id string) *Node {
	return &Node{
		ID:     id,
		HashID: crc32.Checksum([]byte(id), nil),
	}
}

// AddNode : adds node to the ring
func (r *Ring) AddNode(id string) {
	node := NewNode(id)
	r.Nodes = append(r.Nodes, node)

	sort.Slice(r.Nodes, func(i, j int) bool {
		return r.Nodes[i].HashID < r.Nodes[j].HashID
	})
}

// RemoveNode : removes node from the ring
func (r *Ring) RemoveNode(id string) {}

// Get : gets node which is mapped to the key.
func (r *Ring) Get(key string) string {
	return ""
}
