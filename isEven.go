package main

import "fmt"

func main() {
  a := isEven(6)
  fmt.Println(a)
}

func isEven(num int) bool {
  if num % 2 == 0 {
    return true
  }
  return false
}

// true
