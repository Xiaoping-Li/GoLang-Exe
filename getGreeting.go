package main

import (
  "fmt"
)

func main() {
  name := "Tom"
  greeting(name)
}

func greeting(name string) {
  fmt.Printf("hello, %s!", name) 
}
