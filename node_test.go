package main

import (
	"testing"
	"strconv"
	"encoding/hex"
	//"fmt"
)

func Test_NewRandomNodeID(t *testing.T) {
	node := NewRandomNodeID()

	if len(node) != 20 {
		t.Error("NodeID not generated properly: " + strconv.Itoa(len(node)))
	}
}

func Test_NewNodeID(t *testing.T) {
	node1 := NewRandomNodeID()
	nodeStr := hex.EncodeToString(node1[0:IdLength])

	node2 := NewNodeID(nodeStr)

	if node1 != node2 {
		t.Error("Expected nodes to be equal")
	}
}

func Test_Equals(t *testing.T) {
	node1 := NewRandomNodeID()
	nodeStr := node1.String()

	node2 := NewNodeID(nodeStr)

	if !node1.Equals(node2) {
		t.Error("Expected nodes to be equal")
	}
}

func Test_LessTrue(t *testing.T) {
	node1 := NewNodeID("0fda68927f2b2ff836f73578db0fa54c29f7fd82")
	node2 := NewNodeID("0fda68927f2b2ff836f73578db0fa54c29f7fd92")

	if !node1.Less(node2) {
		t.Error("Expected node1 to be less than node2")
	}
}

func Test_LessFalse_Equal(t *testing.T) {
	node1 := NewNodeID("0fda68927f2b2ff836f73578db0fa54c29f7fd92")
	node2 := NewNodeID("0fda68927f2b2ff836f73578db0fa54c29f7fd92")

	if node1.Less(node2) {
		t.Error("Expected nodes to be equal")
	}
}

func Test_Xor(t *testing.T) {
	node1 := NewNodeID("0fda68927f2b2ff836f73578db0fa54c29f7fd82")
	node2 := NewNodeID("0fda68927f2b2ff836f73578db0fa54c29f7fd92")

	xor := node1.Xor(node2)

	expected := "0000000000000000000000000000000000000010"
	if xor.String() != expected {
		t.Error("Expected 0000000000000000000000000000000000000010 got " + xor.String())
	}
}