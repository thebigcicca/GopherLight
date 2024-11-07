package router

import (
	"net/http"
)

type Node struct {
	segment  string
	handler  http.HandlerFunc
	children map[string]*Node
}

func NewNode(segment string) *Node {
	return &Node{
		segment:  segment,
		children: make(map[string]*Node),
	}
}

func (n *Node) AddRoute(path []string, handler http.HandlerFunc) {

	if len(path) == 0 {
		n.handler = handler
		return
	}

	nextSegment := path[0]
	child, exists := n.children[nextSegment]
	if !exists {
		child = NewNode(nextSegment)
		n.children[nextSegment] = child
	}

	child.AddRoute(path[1:], handler)
}

func (n *Node) FindRoute(path []string) (http.HandlerFunc, bool) {
	if len(path) == 0 {
		return n.handler, n.handler != nil
	}

	nextSegment := path[0]
	child, exists := n.children[nextSegment]
	if !exists {
		return nil, false
	}

	return child.FindRoute(path[1:])
}
