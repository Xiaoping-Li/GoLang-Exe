/*
Given 2 words, "getLengthOfTwoWords" returns the sum of their lengths.

*/


package main

import "fmt"

func main() {
  word1 := "Hello"
  word2 := "World"
  getLengthOfTwoWords(word1, word2)
}

func getLengthOfTwoWords(word1, word2 string) {
  var n int
  n = len(word1) + len(word2)
  fmt.Printf("The length of %s and %s is %d", word1, word2, n)
}

// The length of Hello and World is 10 
