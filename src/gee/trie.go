package gee

import (
	"strings"
)

type node struct {
	pattern string 	// the whole path of parts from root to this node
	part string // the part of route the node holds, e.g. :lang, 
				// prefix : means this position should be a param
	children []*node // children nodes, e.g. [doc, tutorial, intro]
	isWild bool // whether to match exactly, if the part has : or * it's true
}

// The first matched child node
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// All the matched child node
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// Insert a route to trie
// Pattern is the whole path
// parts is pattern divided by /
// Height is the current deepth
// Not until the last node, set the pattern of the node
func (n *node) insert(pattern string, parts[]string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	child.insert(pattern, parts, height+1)
}

// Search the left node whose pattern matches parts combined
// If meeting prefix * or using the whole parts, the recursion will quit
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}

	return nil
}