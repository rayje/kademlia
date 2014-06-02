package main

import (
	"encoding/hex"
	"rand"
)

const IdLength = 20

type NodeID [IdLength]byte

// Creates a new NodeID from a hex string
func NewNodeID(data string) (ret NodeID) {
	decoded, _ := hex.DecodeString(data)

	for i := 0; i < IdLength; i++ {
		ret[i] = decoded[i]
	}

	return
}

// Generates a new NodeID from random ints
func NewRandomNodeID() (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = uint8(rand.Intn(256))
	}

	return
}

// Returns a string representation of the NodeID
func (node NodeID) String() string {
	return hex.EncodeToString(node[0:IdLength])
}

// Compares the current NodeID to determine if they are equal.
// Returns true if they are equal, false otherwise.
func (node NodeID) Equals(other NodeID) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other[i] {
			return false
		}
	}
	return true
}

// Determine if this node is less than the other
func (node NodeID) Less(other interface{}) bool {
	for i := 0; i < IdLength; i++ {
		if node[i] != other.(NodeID)[i] {
			return node[i] < other.(NodeID)[i]
		}
	}
	return false
}

// Use an Xor comparison to determin distance
func (node NodeID) Xor(other NodeID) (ret NodeID) {
	for i := 0; i < IdLength; i++ {
		ret[i] = node[i] ^ other[i]
	}
	return
}
