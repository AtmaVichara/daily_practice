package main

type Tree struct {
  Root *Node
}

func TreeInit() *Tree {
  return &Tree {Root: nil}
}

func (t *Tree) Insert(value, data string) error {
  newNode := &Node {Value: value, Data: data}
  if t.Root == nil {
    t.Root = newNode
    return nil
  }

  return t.Root.Insert(value, data)
}

func (t *Tree) Find(s string) string {
  if t.Root == nil {
    return "Root is Nil"
  }

  return t.Root.Find(s)
}
