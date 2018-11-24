/*
Write a function called "isOldEnoughToDrive".

Given a number, in this case an age, "isOldEnoughToDrive" returns whether a person of this given age is old enough to legally drive in the United States.

Notes:* The legal driving age in the United States is 16.

*/


package main

import "fmt"

func main() {
  a := 12
  isOldEnoughToDrive(a)
}

func isOldEnoughToDrive(age int) {
  if age >= 16 {
    fmt.Println("true")
  } else {
    fmt.Println("false")
  }
}
