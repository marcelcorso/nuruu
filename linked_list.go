package nuruu

type ListNode struct {
	key  interface{}
	next *ListNode
}

type LinkedList struct {
	head *ListNode
	tail *ListNode
}

func (ll *LinkedList) add(key interface{}) {
	node := &ListNode{key: key}
	if ll.head == nil {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}
}

func (ll *LinkedList) each(visit func(key interface{}) bool) {
	walker := ll.head
	for walker != nil {
		if (visit(walker.key) == false) {
      break
    }
		walker = walker.next
	}
}

/*
func (ll *LinkedList) member(key interface{}) {
}
*/

