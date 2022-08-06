package main

import "errors"

// Err
var (
	ErrInsufficientMemberCount = errors.New("insufficient member count")
	ErrMemberNotFound          = errors.New("member could not be found in ring")
)

// Hasher is responsible for generating unsigned, 64 bit hash of provided byte slice.
// Hasher should minimize collisions (generating same hash for different byte slice)
// and while performance is also important fast functions are preferable
type Hasher interface {
	Sum64([]byte) uint64
}

// Member interface represents a member in consistent hash in ring
type Member interface {
	String() string
}

// Config represents a structure to control consistent package
type Config struct {
	// Hasher is responsible for generating unsigned 64 bit hash of provided byte slice
	Hasher Hasher
	// Keys are distributed among partitions. Prime numbers are good to distribute keys uniformly.
	// Select a big ParitionCount if you have too many keys
	ParitionCount int
	// Members are replicated on consistent hash ring. This number means that a member how many times repliated on the ring
	ReplicationFactor int
	// Load is used to caculate average load.
	Load float64
}
