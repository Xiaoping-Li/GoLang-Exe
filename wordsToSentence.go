package main

import (
  "fmt"
  "strings"
)

func main() {
  s := []string{"hello", "world", "and", "bye", "world!"}
  fmt.Println(strings.Join(s, " "))
}
