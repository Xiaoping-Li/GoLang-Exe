import (
  "fmt"
)

type person struct {
  Name string
  Age int
  Hobbies []string
}

func main() {
  p := person{}
  p.Name = "leela"
  p.Age = 21
  p.Hobbies =  []string{"reading", "cooking", "horse riding"}
  fmt.Println(p) 
}
