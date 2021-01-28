package main

import "fmt"

func main() {
  a := []int{2,3,4,5,6}
  fmt.Println("hw", len(a))
  fmt.Println(a[:len(a)])
  fmt.Println(a[1:])
  fmt.Println(a[:])
  var b []int
  fmt.Println("b", b)
  b = append(b, 1)
  fmt.Println("b", b)
  b = append(b, 1, 2, 3, 4)
  fmt.Println("b", b)
  for i, v := range a {
    fmt.Println(i,v)
  }
  for _, v := range a {
    fmt.Println(v)
  }
  mm := make(map[string]string)
  mm["hello"] = "world"
  fmt.Println(mm)
  mm = map[string]string{
    "bye bye": "underworld",
  }
  fmt.Println(mm)
  _, ok := mm["hello"]
  if !ok {
    fmt.Println("oh no")
  }
}
