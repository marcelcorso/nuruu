package nuruu

import "testing"

func TestFirstAdd(t *testing.T) {
	ll := &LinkedList{}
	ll.add(55)

	if ll.head.key != 55 {
		t.Errorf("head key is %d, want %d", ll.head.key, 55)
	}

	if ll.tail.key != 55 {
		t.Errorf("tail key is %d, want %d", ll.head.key, 55)
	}
}

func TestTwoAdds(t *testing.T) {
	ll := &LinkedList{}
	ll.add(55)
	ll.add(88)

	if ll.head.key != 55 {
		t.Errorf("head key is %d, want %d", ll.head.key, 55)
	}

	if ll.tail.key != 88 {
		t.Errorf("tail key is %d, want %d", ll.head.key, 88)
	}

	if ll.head.next.key != 88 {
		t.Errorf("head next key is %d, want %d", ll.head.next.key, 88)
	}
}

func TestEach(t *testing.T) {
	ll := &LinkedList{}
	var stuff = []int{1, 2, 4, 8, 16, 32, 64, 128, 77}
	for _, v := range stuff {
		ll.add(v)
	}

	i := 0
	ll.each(func(key interface{}) bool {
		if key != stuff[i] {
			t.Errorf("ll[%d] is %d, want %d", i, key, stuff[i])
		}
		i++
    return true
	})
}

