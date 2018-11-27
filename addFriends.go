package main

import (
  "fmt"
)

type user struct {
  Name string
  Friends []string
}

func main() {
  u := &user{}
  u.Name = "leela"
  u.Friends = []string{}
  u.addFriend("golang")
  u.addFriend("JS")
  fmt.Println(u) 
}

func (u *user) addFriend (f string) {
  u.Friends = append(u.Friends, f)
}
