package main

import "fmt"

type name struct {
  first string
  last string
}

func main() {
  n := name{"leela", "lee"}
  addFullNameProperty(n)
}

func addFullNameProperty(n name) {
  first := n.first
  last := n.last
  fmt.Printf("The full-name is: %s %s", first, last)
}
