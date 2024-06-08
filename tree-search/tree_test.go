package main

import "testing"

func TestEmptyTree(t *testing.T) {
	var node *Node
	result := CheckDuplicateIDs(node)

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestTreeSingleNode(t *testing.T) {
	node := &Node{ value: "1" }
	result := CheckDuplicateIDs(node)

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}

func TestTreeSingleDepthNoDuplicates(t *testing.T) {
	node := &Node{ value: "1", children: []*Node{{ value: "2" }, { value: "3" }} }
	result := CheckDuplicateIDs(node)

	if result != nil {
		t.Errorf("Expected nil, got %v", result)
	}
}


func TestTreeSingleDepthDuplicates(t *testing.T) {
	firstNode := &Node{ value: "2" }
	secondNode := &Node{ value: "1" }
	root := &Node{ value: "1", children: []*Node{firstNode, secondNode} }

	result := CheckDuplicateIDs(root)

	if result != secondNode {
		t.Errorf("Expected %v, got %v", secondNode, result)
	}
}

func TestTreeDuplicatesDeeperLevel(t *testing.T) {
	firstChildFirstNode := &Node{ value: "4" }
	secondChildFirstNode := &Node{ value: "5" }
	firstChildSecondNode := &Node{ value: "2" }
	secondChildSecondNode := &Node{ value: "4" }

	firstNode := &Node{ value: "2" }
	secondNode := &Node{ value: "3" }

	firstNode.children = []*Node{firstChildFirstNode, secondChildFirstNode }
	secondNode.children = []*Node{firstChildSecondNode, secondChildSecondNode }

	root := &Node{ value: "1", children: []*Node{firstNode, secondNode} }

	result := CheckDuplicateIDs(root)

	if result != firstChildSecondNode {
		t.Errorf("Expected %v, got %v", secondNode, result)
	}
}