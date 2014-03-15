package nuruu

type Node struct {
  key int
  next *Node
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (ll *LinkedList) add(key int) {
  node := &Node{key: key}
  if ll.head == nil {
    ll.head = node
    ll.tail = node
  } else {
    ll.tail.next = node
    ll.tail = node
  }
}

func (ll *LinkedList) each(visit func(key int)) {
  walker := ll.head
  for walker != nil {
    visit(walker.key)
    walker = walker.next
  }
}

