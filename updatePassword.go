// Method updatePW will update person's PW
package main

import (
  "fmt"
)

type person struct {
  Email string
  PW string
}

func main() {
  p := &person{}
  p.Email = "leela@test.com"
  p.PW = "123"
  p.updatePW("I love go")
  fmt.Println(p.PW) 
}

func (p *person) updatePW(pw string) {
  p.PW = pw
}
