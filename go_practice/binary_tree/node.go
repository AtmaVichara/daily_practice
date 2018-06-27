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

func (n *Node) findMax(parent *Node) (*Node, *Node) {
  if n == nil {
    return nil, parent
  }
  if n.Right == nil {
    return n, parent
  }
  return n.Right.findMax(n)
}

func (n *Node) replaceNode(parent, replacement *Node) error {
  if n == nil {
    return errors.New("replaceNode() cannot work on node nil!")
  }

  if n == parent.Left {
    parent.Left = replacement
    return nil
  }
  parent.Right = replacement
  return nil
}

func (n *Node) Delete(s string, parent *Node) string {
  if n == nil {
    return "Node is nil"
  }

  switch {
    case s > n.Value:
      return n.Right.Delete(s, n)
    case s < n.Value:
      return n.Left.Delete(s, n)
    default:
      if n.Right == nil && n.Left == nil {
        n.replaceNode(parent, nil)
        return "Deleted " + s
      }
      if n.Left == nil {
        n.replaceNode(parent, n.Right)
        return "Deleted " + s
      }
      if n.Right == nil {
        n.replaceNode(parent, n.Left)
        return "Deleted " + s
      }
      replacement, replParent := n.Left.findMax(parent)

      n.Value = replacement.Value
      n.Data = replacement.Data

      return replacement.Delete(replacement.Value, replParent)
  }

}
