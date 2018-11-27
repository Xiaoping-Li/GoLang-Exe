package main

import (
  "fmt"
)

func main() {
  a := []int{1,2,3}
  fmt.Println(incrementByOne(a))
  
}

func incrementByOne(a []int) []int {
  for i := range a {
    a[i]++
  }
  return a
}
