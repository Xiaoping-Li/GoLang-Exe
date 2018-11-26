// append item(s) into the end of slice
package main

import (
  "fmt"
)

func main() {
  s := []string{"hello", "world"}
  a := append(s, "append")
  fmt.Println(a)
}
// [hello world append]



// prepend item(s) into the front of slice
package main

import (
  "fmt"
)

func main() {
  s := []string{"hello", "world"}
  a := append([]string{"prepend"}, s...)
  fmt.Println(a)
}
// [prepend hello world]
