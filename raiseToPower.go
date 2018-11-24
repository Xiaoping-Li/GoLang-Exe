/*
Raise num to whatever power is passed in as the exponent.
*/


package main

import (
  "fmt"
  "math"
)

func main() {
  var x, y float64
  x = 2.5
  y = 3
  rtn := raiseToPower(x, y) 
  fmt.Println(rtn) 
}

func raiseToPower(x, y float64) float64 {
  return math.Pow(x, y)
}
