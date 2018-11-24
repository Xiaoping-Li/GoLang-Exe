package main

import "fmt"

func main() {
  var n float32
  n = 3.01
  rtn := multiplyByTen(n)
  fmt.Println(rtn)
}

func multiplyByTen(num float32) float32 {
  return num * 10
}
