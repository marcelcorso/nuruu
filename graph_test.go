package nuruu

import (
  "fmt"
  "testing"
  "strconv"
)

func TestBasicTree(t *testing.T) {
  g := NewGraph()
  g.Add("a", "b")
  g.Add("a", "c")
  g.Add("b", "c")
  g.Add("b", "d")
  g.Add("c", "d")

  if g.Adjacent("a", "b") == false {
    t.Errorf("should have %d-%d edge", "a", "b")
  }

  if g.Adjacent("b", "a") == false {
    t.Errorf("should have %d-%d edge", "b", "a")
  }

  if g.Adjacent("b", "d") == false {
    t.Errorf("should have %d-%d edge", "b", "d")
  }

  if g.Adjacent("z", "a") == true {
    t.Errorf("should NOT have %d-%d edge", "z", "a")
  }

  if g.Adjacent("a", "d") == true {
    t.Errorf("should NOT have %d-%d edge", "c", "a")
  }


  if g.Adjacent("a", "a") == true {
    t.Errorf("a-a edges are not allowed")
  }

  g.Remove("a", "b")

  if g.Adjacent("a", "b") == true {
    t.Errorf("a-b should be gone")
  }
}
