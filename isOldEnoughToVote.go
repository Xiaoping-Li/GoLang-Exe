/*
Write a function called "isOldEnoughToVote".

Given a number, in this case an age, 'isOldEnoughToVote' returns whether a person of this given age is old enough to legally vote in the United States.

Notes:* The legal voting age in the United States is 18.
*/


package main

import "fmt"

func main() {
  a := 20
  isOldEnoughToVote(a)
}

func isOldEnoughToVote(age int) {
  if age >= 18 {
    fmt.Println("true")
  } else {
    fmt.Println("false")
  }
}
