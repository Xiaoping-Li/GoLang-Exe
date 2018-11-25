/*
testScores is an array.  Iterate over testScores and compute and return the average
*/


package main

import (
  "fmt"
)

func main() {
  s := []float64{1.1,2.2,3.3,4,5}
  rtn := averageTestScore(s)
  fmt.Println(rtn)
}

func averageTestScore(nums []float64) float64 {
  sum := 0.0
  n := float64(len(nums))
  for _, v := range nums {
    sum += v
  }
  return sum / n 
}
