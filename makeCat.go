// Make cat struct with Name
// Add meow method to cat struct
package main

import (
  "fmt"
)

type cat struct {
  Name string
}

func main() {
  newCat := cat{"Snow ball"}
  newCat.meow()
}

func (c cat) meow() {
  fmt.Printf("%s meowed!!!", c.Name)
}
