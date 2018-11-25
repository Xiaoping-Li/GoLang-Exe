/*
Return the factorial of num.
The factorial of a number is the product of each number between 1 and n.

Input Example:

5
10
2

Output Example:

120 // 5 * 4 * 3 * 2 * 1
3628800 // 10 * 9 * 8 * 7 * 6 * 5 * 4 * 3 * 2 * 1
2 // 2 * 1
*/


package main

import (
  "fmt"
)

func main() {
  num := 1
  rtn := nFactorial(num)
  fmt.Println(rtn)
}

func nFactorial(num int) int {
  rtn := 1
  for i := 1; i <= num; i++ {
    rtn *= i
  }
  return rtn
}
