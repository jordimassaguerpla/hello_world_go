package main

import (
  "fmt"
  "runtime"
)

type Vertex struct {
  X int
  Y int
}

func main() {
  defer fmt.Println("Bye bye")
  for i := 0; i < 10; i++ {
    fmt.Printf("hello world!\n")
    if i > 1 {
      break;
    }
  }
  ooo := "AAA"
  switch os := runtime.GOOS; os {
  case "linux":
    fmt.Printf("linux")
  case "windows":
    fmt.Printf("oh no")
  default:
    fmt.Printf("dont know")
  }
  p_os := &ooo
  fmt.Printf(*p_os)
  *p_os = "oh no"
  fmt.Printf(ooo)
  v := Vertex{1, 2}
  fmt.Println(v.X)
  p := & v
  fmt.Println(p.X)
  a := []string{"hello","world"}
  var b [2]string
  b[0] = "bye"
  b[1] = "underworld"
  fmt.Println(a, b, len(a))
}
