package main

import (
	"fmt"

	"github.com/buraksezer/consistent"
	"github.com/cespare/xxhash"
)

type myMember string

func (m myMember) String() string {
	return string(m)
}

type hasher struct{}

func (h hasher) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}

func main() {
	cfg := consistent.Config{
		Hasher:            hasher{},
		PartitionCount:    7,
		ReplicationFactor: 20,
		Load:              1.25,
	}

	c := consistent.New(nil, cfg)

	node1 := myMember("node1.olricmq.com")
	c.Add(node1)
	node2 := myMember("node2.olricmq.com")
	c.Add(node2)

	key := []byte("my-keyssss")
	// calculates partition id for the given key
	// partID := hash(key) % partitionCount
	// the partitions is already distributed among members by Add function.
	owner := c.LocateKey(key)
	fmt.Println(owner.String())
	// Prints node2.olricmq.com
}
