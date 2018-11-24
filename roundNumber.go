/*
Round num and return it.
*/

package main

import (
  "fmt"
  "math"
)

func main() {
  var x float64
  x = 3.4
  rtn := math.Round(x)
  fmt.Println(rtn)
}
