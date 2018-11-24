/*
Return a string of all the even numbers between 0 and num inclusive.

Example Input:

7
12
25

Example Output:

0246
024681012
024681012141618202224

Note: 0 is an even number
*/



package main

import (
  "fmt"
  "strconv"
)

func main() {
  num := 11
  rtn := stringOfEvens(num)
  fmt.Println(rtn)
}

func stringOfEvens(num int) string {
  rtn := ""
  for i := 0; i <= num; i += 2 {
    n := strconv.Itoa(i)
    rtn += n
  }
  return rtn
}
