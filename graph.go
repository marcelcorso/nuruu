package nuruu

import (
  "fmt"
)

type Edge struct {
  row, column string
  // a sorted linked list happens here. it could be a map too. nah. later maybe.
  Next *Edge
}

type Graph struct {
	rows map[string]*Edge
}

func NewGraph() *Graph {
    g := new(Graph)
    g.rows = make(map[string]*Edge)
    return g
}

func (g *Graph) Add(row string, column string) bool {
  if row == column {
    return false
  }
  // be sure row < column
  if row < column {
    row, column = column, row
  }

  edge := Edge{row, column, nil}

  var previous *Edge = nil
  p := g.rows[row]

  for ;; {
    if p == nil || p.column > column {
      if previous == nil {
        g.rows[row] = &edge
        return true
      } else {
        previous.Next = &edge
        return true
      }
    }
    previous = p
    p = p.Next
  }

}

func (g *Graph) Adjacent(row string, column string) bool {
  if row == column {
    return false
  }
  // be sure row < column
  if row < column {
    row, column = column, row
  }
  for p := g.rows[row]; (p != nil && p.column <= column); p = p.Next {
    if p.column == column {
      return true
    }
  }
  return false
}

