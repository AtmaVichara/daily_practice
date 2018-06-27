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

func (t *Tree) Delete(s string) string {
  if t.Root == nil {
    return "Root is Nil"
  }

  fakeParent := &Node{Right: t.Root}
  message := t.Root.Delete(s, fakeParent)

  if fakeParent.Right == nil {
    t.Root = nil
    return "Deleted Element. Tree is empty."
  }

  return message
}

func (t *Tree) Traverse (n *Node, f func(*Node)) {
  if n == nil {
    return
  }
  t.Traverse(n.Left, f)
  f(n)
  t.Traverse(n.Right, f)
}
