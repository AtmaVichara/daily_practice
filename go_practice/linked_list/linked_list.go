package main

import "fmt"

type Node struct {
  next *Node

  val interface{}
}

type List struct {
  head *Node

  count int
}

func (L *List) Prepend(val interface{}) {
  node := &Node{
    next: L.head,
    val: val,
  }
  L.head = node
  L.count++
}

func (L *List) Append(val interface{}) {
  node := &Node{
    val: val,
  }

  if L.head == nil {
    L.head = node
  } else {
    l := L.head
    for l.next != nil {
      l = l.next
    }
    l.next = node
  }
  L.count++
}

func (L *List) Insert(position int, val interface{}) {
  l := L.head
  for i := 1; i < position; i++ {
    l = l.next
  }
  node := &Node{
    next: l.next,
    val: val,
  }
  l.next = node
  L.count++
}

func (L *List) Pop() interface{} {
  if L.head.next == nil {
    holder := L.head.val
    L.head = nil
    return holder
  } else {
    l := L.head
    for l.next.next != nil {
      l = l.next
    }
    value := l.next.val
    l.next = nil
    L.count--
    return value
  }
}

func (L *List) Shift() interface{} {
  value := L.head.val
  L.head = L.head.next
  L.count--
  return value
}

func (L *List) Count() int{
  return L.count
}

func (L *List) Show() {
  list := L.head
  for list != nil {
    fmt.Printf("%v -> ", list.val)
    list = list.next
  }
  fmt.Println()
}

func main() {
  l := List{}
  l.Append(10)
  l.Append(10)
  l.Append(10)

  l.Show()
  l.Prepend(12)
  l.Prepend(13)
  l.Prepend(14)

  l.Show()

  l.Append(15)
  l.Show()

  l.Insert(3, 20)
  l.Insert(3, 22)
  l.Show()

  fmt.Println(l.Pop())
  l.Show()

  fmt.Println(l.Shift())
  l.Show()

  fmt.Printf("count: %v\n", l.Count())
}
