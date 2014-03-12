package main

import (
  "fmt"
)

type Node struct {
  value int
  next *Node
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (ll *LinkedList) add(value int) {
  node := &Node{value: value}
  if ll.head == nil {
    ll.head = node
    ll.tail = node
  } else {
    ll.tail.next = node
    ll.tail = node
  }
}

func (ll *LinkedList) print() {
  if ll.head == nil {
    fmt.Printf("<< empty >>")
  } else {
    walker := ll.head
    for walker != nil {
      fmt.Printf("%d -> ", walker.value)
      walker = walker.next
    }
  }
}

func main() {
  ll := &LinkedList{}
  for i := 0; i < 10; i++ {
    ll.add(i)
  }

  ll.print()
  fmt.Printf("\n")
}


