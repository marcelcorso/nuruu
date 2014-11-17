package nuruu

import (
  "testing"
)

func TestBasicGraph(t *testing.T) {
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

func TestBasicLists(t *testing.T) {
  g := NewGraph()
  givenEdges := [][]string{{"a", "b"}, {"a", "c"}, {"b", "c"}, {"b", "d"}, {"c", "d"}}
  for _, e := range(givenEdges) {
    g.Add(e[0], e[1])
  }

  vertices := g.Vertices()
  expectedVertices := []string{"a", "b", "c", "d"}
  if (len(vertices) != len(expectedVertices)) {
    t.Errorf("The lenght of the vertices slice is wrong.")
  } else {
    for _, v := range(expectedVertices) {
      found = false
      for _, u := range(vertices) {
        if v == u {
          found = true
          break
        }
      }
      if found == false {
        t.Errorf("Can't find expected vertice '" + v + "' in the vertices slice.")
      }
    }
  }

  edges := g.Edges()
  if (len(edges) != len(givenEdges)) {
    t.Errorf("The lenght of the edges slice is wrong.")
  } else {
    for _, e := range(givenEdges) {
      if e[0] > e[1] {
        e[0], e[1] = e[1], e[0]
      }
      found = false
      for _, s := range(edges) {
        if s[0] > s[1] {
          s[0], s[1] = s[1], s[0]
        }

        if e[0] == s[0] && e[1] == s[1] {
          found = true
        }

      }
      if found == false {
        t.Errorf("Can't find expected edge " + e[0] + "-" + e[1] + " in the edges slice.")
      }
    }
  }
}

func TestGraphDegree(t *testing.T) {
  g := NewGraph()
  g.Add("a", "b")
  g.Add("a", "c")
  g.Add("b", "c")
  g.Add("b", "d")
  g.Add("c", "d")

  g.Add("z", "k")

  if g.Degree("a") != 2 {
    t.Errorf("Degree of 'a' should be 2")
  }

  if g.Degree("c") != 3 {
    t.Errorf("Degree of 'c' should be 3")
  }

  g.Remove("a", "c")

  if g.Degree("c") != 2 {
    t.Errorf("Degree of 'c' should be 2 after removing an edge to it")
  }

  if g.Degree("z") != 1 {
    t.Errorf("Degree of 'z' should be 1")
  }

}

func TestGraphVerticeOps(t *testing.T)
  // TODO 
  // ./daringanga.go:49: h.RemoveVertices undefined (type *nuruu.Graph has no field or method RemoveVertices)
  t.Errorf("TestGraphVerticeOps should be inplemented")

}
