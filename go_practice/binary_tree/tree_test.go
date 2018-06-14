package main

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func TestNewTree(t *testing.T) {
  tree := TreeInit()

  assert.Nil(t, tree.Root)
}

func TestTreeInsert(t *testing.T) {
  tree := TreeInit()

  tree.Insert("d", "delta")
  assert.Equal(t, tree.Root.Value, "d", "they should be equal")
  assert.Equal(t, tree.Root.Data, "delta", "they should be equal")

  tree.Insert("a", "alpha")
  assert.Equal(t, tree.Root.Left.Value, "a", "they should be equal")
  assert.Equal(t, tree.Root.Left.Data, "alpha", "they should be equal")

  tree.Insert("e", "echo")
  assert.Equal(t, tree.Root.Right.Value, "e", "they should be equal")
  assert.Equal(t, tree.Root.Right.Data, "echo", "they should be equal")


  tree.Insert("b", "bravo")
  nextLeaf := tree.Root.Left
  assert.Equal(t, nextLeaf.Right.Value, "b", "they should be equal")
  assert.Equal(t, nextLeaf.Right.Data, "bravo", "they should be equal")
}

func TestTreeFind(t *testing.T) {
  tree := TreeInit()
  tree.Insert("a", "alpha")
  tree.Insert("e", "echo")
  tree.Insert("b", "bravo")
  tree.Insert("d", "delta")

  assert.Equal(t, tree.Find("a"), "alpha", "Find should return alpha for 'a'")
  assert.Equal(t, tree.Find("d"), "delta", "Find should return delta for 'd'")
  assert.Equal(t, tree.Find("t"), "Tree does not contain value", "they should be equal")
}

func TestTreeDelete(t *testing.T) {
  tree := TreeInit()
  tree.Insert("a", "alpha")
  tree.Insert("e", "echo")
  tree.Insert("b", "bravo")
  tree.Insert("d", "delta")

  tree.Delete("d")
  assert.Equal(t, tree.Find("d"), "Tree does not contain value", "they should be equal")
}
