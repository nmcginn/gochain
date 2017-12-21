package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestTree(t *testing.T) {
	merkle_tree := build_test_tree()
	if merkle_tree.Value != "test" {
		t.Error("root value for test tree was incorrect")
	}
	if merkle_tree.Left.Value != "xxx" || merkle_tree.Right.Value != "ccc" {
		t.Error("tree is not balanced correctly")
	}
}
