/*
If num is divisible by 3 return 'fizz'.
If num is divisible by 5 return 'buzz'.
If num is divisible by 3 & 5 return 'fizzbuzz'.
Otherwise return num.
*/


package main

import (
  "fmt"
)

func main() {
  num := 30
  fizzBuzz(num)
}

func fizzBuzz(num int) {
  switch {
    case num % 15 == 0:
      fmt.Printf("fizzbuzz")
    case num % 3 == 0:
      fmt.Printf("fizz")
    case num % 5 == 0:
      fmt.Printf("buzz")
    default:
      fmt.Println(num)
  }
}
