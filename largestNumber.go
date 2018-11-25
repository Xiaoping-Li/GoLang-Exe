/*
numbers is an array of integers.  Find and return the largest integer.

Input Example:

[1, 2, 3, 4, 5]
[100, 0, -100]
[5, 5, 5]

Output Example:

5
100
5
*/


package main

import (
  "fmt"
)

func main() {
  s := []int{1,2,3,4,5}
  largestNumber(s)
}

func largestNumber(nums []int) {
  large := nums[0]
  for _, v := range nums {
    if v > large {
      large = v
    }
  }
  fmt.Println(large)
}
