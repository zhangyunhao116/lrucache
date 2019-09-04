package lrucache

import (
	"testing"
)

func TestDeepCopyNode(t *testing.T) {
	mockPrevNode := node{}
	mockNextNode := node{}
	node := node{key: "111", value: "222", prev: &mockPrevNode, next: &mockNextNode}
	newNode := deepCopyNode(node)
	if &node == &newNode {
		t.Error("error same address")
	}

	if node.key != newNode.key || node.value != newNode.value || node.prev != newNode.prev || node.next != newNode.next {
		t.Error("error content")
	}

}

func TestNewLRUCache(t *testing.T) {
	l := New(3)

	// Simple set
	replaceKey := l.Set(1, 15)
	if replaceKey {
		t.Error("error replace while set")
	}

	// Simple Get
	v, ok := l.Get(1)
	if v != 15 || !ok {
		t.Error("error get exists key")
	}

	// Get non-existent key
	v, ok = l.Get(2)
	if v != nil || ok {
		t.Error("error get non-existent key")
	}

	// Get eliminated key
	l.Set(2, 20)
	l.Set(3, 30)
	l.Set(4, 40)

	v, ok = l.Get(1)
	if v != nil || ok {
		t.Error("error get eliminated key")
	}

	// Get sort
	// Now the key in cache is (4,3,2)
	v, ok = l.Get(2) // Cache : (2,4,3)
	if v != 20 || !ok {
		t.Error("error get exists key")
	}

	l.Set(5, 50) // Now should be (5,2,4)
	v,_ = l.Get(2)
	v, ok = l.Get(3)
	if v != nil || ok {
		t.Error("error get sort")
	}

}
