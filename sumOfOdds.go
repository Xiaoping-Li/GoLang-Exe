/*
Sum all of the odd numbers together between 1 and num inclusive.
Assume num is a positive integer.

Input Example:

5
15
2

Output Example:

9 // 5 + 3 + 1
64 // 15 + 13 + 11 + 9 + 7 + 5 + 3 + 1
1 // 1
*/


package main

import (
  "fmt"
)

func main() {
  num := 15
  rtn := sumOfOdds(num)
  fmt.Println(rtn)
}

func sumOfOdds(num int) int {
  rtn := 0
  for i := 1; i <= num; i += 2 {
    rtn += i
  }
  return rtn
}
