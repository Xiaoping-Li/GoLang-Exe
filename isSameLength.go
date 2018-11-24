package main

import "fmt"

func main() {
  rtn := isSameLength("ni", "hao")
  fmt.Println(rtn)
}

func isSameLength(word1, word2 string) bool {
  a := len(word1)
  b := len(word2)

  if a == b {
    return true
  }
  return false
}
