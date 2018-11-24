/*
return true if x and y are the same.  Otherwise return false.
*/


package main

import (
  "fmt"
  "reflect"
)

func main() {
  x := 2
  y := "2"
  rtn := areEqual(x, y)
  fmt.Println(rtn)
}

func areEqual(x, y interface{}) bool {
  xT := reflect.TypeOf(x).String() 
  yT := reflect.TypeOf(y).String()

  if xT == yT && x == y {
    return true
  } else {
    return false
  }
}
