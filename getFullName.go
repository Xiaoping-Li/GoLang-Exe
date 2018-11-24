/*
Given a first and a last name, "getFullName" returns a single string with the given first and last names separated by a single space.
*/


package main

import "fmt"

func main() {
  first := "leela"
  last := "lee"
  getFullName(first, last)
}

func getFullName(first, last string) {
  fmt.Printf("%s %s", first, last)
}
