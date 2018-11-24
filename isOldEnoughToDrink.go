/*
Write a function called "isOldEnoughToDrink".

Given a number, in this case an age, "isOldEnoughToDrink" returns whether a person of this given age is old enough to legally drink in the United States.

Notes:* The legal drinking age in the United States is 21.

*/

package main

import "fmt"

func main() {
  a := 12
  isOldEnoughToDrink(a)
}

func isOldEnoughToDrink(age int) {
  if age >= 21 {
    fmt.Println("true")
  } else {
    fmt.Println("false")
  }
}
