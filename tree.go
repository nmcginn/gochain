package main

import (
	"strings"
)

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func add_node(root *Node, value string) *Node {
	side := strings.Compare(root.Value, value)
	if side <= 0 {
		if root.Left == nil {
			new_node := Node{
				Value: value,
			}
			root.Left = &new_node
			return root
		} else {
			add_node(root.Left, value)
		}
	} else {
		if root.Right == nil {
			new_node := Node{
				Value: value,
			}
			root.Right = &new_node
			return root
		} else {
			add_node(root.Right, value)
		}
	}
	return root
}

func build_test_tree() *Node {
	root := Node{
		Value: "test",
		Left:  nil,
		Right: nil,
	}
	add_node(&root, "ccc")
	add_node(&root, "xxx")
	return &root
}
