package main

import (
  "errors"
)

type Node struct {
  Value string
  Data string
  Left *Node
  Right *Node
}

func (n *Node) Insert(value, data string) error {
  if n == nil {
    return errors.New("Cannot Insert value into nil node.")
  }

  newNode := &Node {Value: value, Data: data}

  switch {
    case value == n.Value:
      return nil
    case value > n.Value:
      if n.Right == nil {
        n.Right = newNode
      } else {
        n.Right.Insert(value, data)
      }
    case value < n.Value:
      if n.Left == nil {
        n.Left = newNode
      } else {
        n.Left.Insert(value, data)
      }
  }
  return nil
}

func (n *Node) Find(s string) string {
  if n == nil {
    return "Tree does not contain value"
  }

  switch {
    case s == n.Value:
      return n.Data
    case s > n.Value:
      return n.Right.Find(s)
    default:
      return n.Left.Find(s)
  }
}
