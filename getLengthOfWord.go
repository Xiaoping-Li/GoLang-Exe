/*
Given a word, "getLengthOfWord" returns the length of the given word.
*/


package main

import "fmt"

func main() {
  word := "Hello"
  getLengthOfWord(word)
}

func getLengthOfWord(word string) {
  var n int
  n = len(word)
  fmt.Printf("The length of %s is %d", word, n)
}


//The length of Hello is 5 
